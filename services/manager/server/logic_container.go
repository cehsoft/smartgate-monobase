package server

import (
	"context"

	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	q "github.com/doug-martin/goqu/v9"

	"github.com/init-tech-solution/service-spitc-stream/services/lib/uuid"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/model"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req)

	trackingType, trackingSession := extractTrackingMetaFromImgURL(req.ImageURL)
	trackingId := new(int)
	if trackingSession != "" {
		id, err := svc.db.ContainerTracking.Create().
			SetSessionID(trackingSession).
			SetContainerID(req.ContainerID).
			OnConflict(
				sql.ConflictColumns(containertracking.FieldSessionID),
			).UpdateNewValues().ID(ctx)
		if err != nil {
			if !strings.Contains(err.Error(), "no rows in result set") {
				return nil, err
			}
		}
		log.Println("tracking_id", id)
		trackingId = &id
	}

	camId, err := svc.db.CamSetting.Query().Where(camsetting.NameEQ(req.CamName)).FirstID(ctx)
	if err != nil {
		return nil, err
	}

	builder := svc.db.ContainerTrackingSuggestion.Create().
		SetResult(req.Result).
		SetImageURL(req.ImageURL).
		SetTrackingType(trackingType).
		SetNillableTrackingID(trackingId).
		SetScore(req.Score).SetCamID(camId)

	if trackingType == "contID" {
		splittedIDs := splitContainerID(req.ContainerID)

		if splittedIDs == nil {
			return nil, fmt.Errorf("result contID is not valid")
		}

		builder.
			SetContainerID(req.Result). // ! DEPRECATED in favor of Result
			SetBic(splittedIDs[0]).
			SetSerial(splittedIDs[1]).
			SetChecksum(splittedIDs[2])
	}

	suggestRrd, err := builder.
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Send to subscribed clients
	for uuid, trackingC := range svc.clients {
		go func(uuid string, trackingC chan *ent.ContainerTrackingSuggestion) {
			timeout := time.NewTimer(10 * time.Second)

			select {
			case trackingC <- suggestRrd:
				log.Println(">>>> sent:", uuid)
			case <-timeout.C:
				log.Println("!!! timeout:", uuid)
			}
		}(uuid, trackingC)
	}

	return &mygrpc.ResEmpty{}, nil
}

func (svc *Server) ListContainerOCRs(ctx context.Context, req *mygrpc.ReqListContainerOCRs) (*mygrpc.ResListContainerOCRs, error) {
	offset, limit := paging(req.Paging)

	camIDs, err := svc.db.CamSetting.Query().Where(camsetting.LaneID(int(req.LaneID))).IDs(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("ListContainerOCRs cams", camIDs)

	if len(camIDs) == 0 {
		return &mygrpc.ResListContainerOCRs{}, nil
	}

	query, _, err := q.Select("bic", "cam_id", "checksum",
		"container_id", "created_at", "id",
		"image_url", "result", "score", "serial", "tracking_id", "tracking_type", q.COUNT("*").Over(q.W()).As("full_count")).
		From(
			q.Select("*", q.ROW_NUMBER().Over(q.W().PartitionBy("tracking_id").OrderBy(q.I("score").Desc())).As("rank")).
				From("container_tracking_suggestions")).
		As("_").
		Where(q.C("rank").Eq(1), q.C("cam_id").In(camIDs)).
		Order(q.I("created_at").Desc()).
		Offset(uint(offset)).
		Limit(uint(limit)).
		ToSQL()
	if err != nil {
		return nil, err
	}

	// ocrs := []*ent.ContainerTrackingSuggestion{}
	ocrs := []*model.ContainerTrackingSuggestion{}
	err = svc.rawdb.Select(&ocrs, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	total := 0
	resOCRs := []*mygrpc.ContainerOCR{}
	for _, s := range ocrs {
		ocr := &mygrpc.ContainerOCR{
			ID:              int32(s.ID),
			Score:           s.Score,
			ImageURL:        s.ImageURL,
			ContainerID:     s.ContainerID,
			Result:          s.Result,
			BIC:             s.Bic,
			Serial:          s.Serial,
			Checksum:        s.Checksum,
			TrackingType:    s.TrackingType,
			TrackingSession: fmt.Sprintf("%d", s.TrackingID),
			CreatedAt:       int32(s.CreatedAt.Unix()),
		}

		resOCRs = append(resOCRs, ocr)
		total = s.FullCount
	}

	return &mygrpc.ResListContainerOCRs{
		Total: int32(total),
		Ocrs:  resOCRs,
	}, nil
}

func (svc *Server) PullMLResult(req *mygrpc.ReqPullMLResult, resp mygrpc.MyGRPC_PullMLResultServer) error {
	// Create client UUID
	clientId := uuid.UUID(0)

	// Create client's data channel
	svc.clients[clientId] = make(chan *ent.ContainerTrackingSuggestion, 10)
	defer func() {
		delete(svc.clients, clientId)
	}()

	ctx := context.Background()

	camIDs, err := svc.db.CamSetting.Query().
		Order(ent.Asc(camsetting.FieldID)).
		Where(camsetting.LaneID(int(req.LaneID))).
		IDs(ctx)
	if err != nil {
		return err
	}

	// Send to browser
	for suggest := range svc.clients[clientId] {
		i := sort.SearchInts(camIDs, suggest.CamID)
		if i < len(camIDs) && camIDs[i] == suggest.CamID {
			err := resp.Send(&mygrpc.ResMLResult{
				SuggestID:   int32(suggest.ID),
				ContainerID: suggest.ContainerID,
				ImageURL:    suggest.ImageURL,
				Score:       suggest.Score,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}
