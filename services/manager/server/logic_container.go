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
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertrackingsuggestion"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/predicate"
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

func (svc *Server) MigrateVirtualSessionsForOCRs(ctx context.Context) error {
	ocrs, err := svc.db.ContainerTrackingSuggestion.Query().
		Where(containertrackingsuggestion.CreatedAtGT(time.Date(2021, time.October, 20, 0, 0, 0, 0, time.UTC))).
		Order(ent.Asc(containertrackingsuggestion.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return err
	}

	virtuals, err := createVirtualSessions(ocrs)
	if err != nil {
		return err
	}

	vss := []string{}
	for vssUUID, _ := range virtuals {
		vss = append(vss, vssUUID)
	}

	tx, err := svc.db.Tx(ctx)
	if err != nil {
		return err
	}

	creatingVss := make([]*ent.ContainerTrackingCreate, len(vss))
	fmt.Println("migrate: total creating virtual sessions:", len(creatingVss))
	for i, name := range vss {
		creatingVss[i] = tx.ContainerTracking.Create().SetSessionID(name)
	}

	createdVss, err := tx.ContainerTracking.CreateBulk(creatingVss...).Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	// j, _ := json.Marshal(virtuals)
	// jj := bytes.NewBuffer(nil)
	// json.Indent(jj, j, "", "\t")
	// fmt.Println(jj.String())

	for _, createdVs := range createdVss {
		sugests := []*ent.ContainerTrackingSuggestion{}

		sugests = append(sugests, virtuals[createdVs.SessionID]["contID"]...)
		sugests = append(sugests, virtuals[createdVs.SessionID]["vehicleID"]...)
		sugests = append(sugests, virtuals[createdVs.SessionID]["romoocID"]...)

		sugestIDs := []predicate.ContainerTrackingSuggestion{}
		for _, s := range sugests {
			sugestIDs = append(sugestIDs, containertrackingsuggestion.ID(s.ID))
		}
		fmt.Printf("migrate: bind %d results to session %s\n", len(sugests), createdVs.SessionID)

		_, err := tx.ContainerTrackingSuggestion.Update().Where(containertrackingsuggestion.Or(
			sugestIDs...,
		)).SetTrackingID(createdVs.ID).Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	_, err = tx.ContainerTracking.Delete().Where(containertracking.Not(containertracking.HasSuggestions())).Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
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
		"image_url", "result", "score", "serial", "tracking_id", "tracking_type", "is_valid",
		q.COUNT("*").Over(q.W()).As("full_count"),
	).
		From(
			q.Select("*", q.ROW_NUMBER().Over(q.W().PartitionBy("tracking_id", "tracking_type").
				OrderBy(q.I("score").Desc())).As("rank")).
				From("container_tracking_suggestions"),
		).
		As("_").
		Where(q.C("cam_id").In(camIDs)).
		// Where(q.C("rank").Eq(1), q.C("cam_id").In(camIDs)).
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
			IsValid:         s.IsValid,
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

func (svc *Server) ValidateOCR(ctx context.Context, req *mygrpc.ReqValidateOCR) (*mygrpc.ResEmpty, error) {
	_, err := svc.db.ContainerTrackingSuggestion.UpdateOneID(int(req.OCRID)).
		SetIsValid(req.Valid).Save(ctx)

	if err != nil {
		return nil, err
	}

	return &mygrpc.ResEmpty{}, nil
}
