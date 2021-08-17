package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/init-tech-solution/service-spitc-stream/services/lib/uuid"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertrackingsuggestion"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/model"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

type Server struct {
	db    *ent.Client
	rawdb *sql.DB

	clients map[string]chan *ContainerTrackingSuggestion

	mygrpc.UnimplementedMyGRPCServer
}

type ContainerTrackingSuggestion struct {
	ID int32
	model.ContainerTracking
}

func CreateServer(db *ent.Client, rawdb *sql.DB) *Server {
	return &Server{
		db:      db,
		rawdb:   rawdb,
		clients: map[string]chan *ContainerTrackingSuggestion{},
	}
}

func paging(req *mygrpc.ReqEmpty) (int, int) {
	offset := 0
	limit := 50

	if req.Limit != nil {
		limit = int(*req.Limit)
	}

	if req.Offset != nil {
		offset = int(*req.Offset) * limit
	}

	return offset, limit
}

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req)

	splittedIDs := splitContainerID(req.ContainerID)

	if splittedIDs == nil {
		return nil, fmt.Errorf("result is not valid")
	}

	suggestRrd, err := svc.db.ContainerTrackingSuggestion.Create().
		SetContainerID(req.ContainerID).
		SetBic(splittedIDs[0]).
		SetSerial(splittedIDs[1]).
		SetChecksum(splittedIDs[2]).
		SetImageURL(req.ImageURL).
		SetScore(req.Score).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	suggest := ContainerTrackingSuggestion{
		ID: int32(suggestRrd.ID),
		ContainerTracking: model.ContainerTracking{
			ContainerID: req.ContainerID,
			ImageURL:    req.ImageURL,
			Score:       req.Score,
		},
	}

	// Send to subscribed clients
	for uuid, trackingC := range svc.clients {
		go func(uuid string, trackingC chan *ContainerTrackingSuggestion) {
			timeout := time.NewTimer(10 * time.Second)

			select {
			case trackingC <- &suggest:
				log.Println(">>>> sent:", uuid)
			case <-timeout.C:
				log.Println("!!! timeout:", uuid)
			}
		}(uuid, trackingC)
	}

	return &mygrpc.ResEmpty{}, nil
}

func (svc *Server) ListContainerTrackings(ctx context.Context, req *mygrpc.ReqEmpty) (*mygrpc.ResListContainerTrackings, error) {
	offset, limit := paging(req)

	total, err := svc.db.ContainerTracking.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	tracks, err := svc.db.ContainerTracking.Query().WithSuggestions().
		Where(containertracking.Not(containertracking.HasSuggestions())).
		Order(ent.Desc(containertracking.FieldCreatedAt)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	// rows, err := svc.rawdb.Query(`
	// WITH dedup_table AS (
	// 	SELECT chunk_time, s.id, s.container_id, s.image_url, s.score, s.bic, s.serial, s.checksum, s.created_at,
	// 		ROW_NUMBER() OVER(PARTITION BY s.serial, s.score ORDER BY id asc) AS row_num
	// 		FROM container_tracking_suggestions AS s
	// 		JOIN (
	// 			SELECT count(*), max(score), serial,
	// 					date_trunc('hour', created_at) +
	// 					(((date_part('minute', created_at)::integer / 10::integer) * 10::integer)
	// 					|| ' minutes')::interval AS chunk_time
	// 			FROM container_tracking_suggestions
	// 			GROUP BY serial, chunk_time
	// 		) AS c
	// 		ON s.score = c.max AND s.serial = c.serial
	// 		ORDER BY created_at desc
	// ) SELECT id, container_id, image_url, score, bic, serial, checksum, created_at
	// 	FROM dedup_table
	// 	WHERE row_num = 1
	// 	LIMIT 100;
	// `)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// suggests, err := svc.db.ContainerTrackingSuggestion.Query().
	// 	Order(ent.Desc(containertrackingsuggestion.FieldCreatedAt)).
	// 	Limit(100).
	// 	All(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	resTrackings := []*mygrpc.ContainerTracking{}

	// for _, s := range suggests {
	// 	splittedIDs := splitContainerID(s.ContainerID)

	// 	if s.Bic != splittedIDs[0] {
	// 		s.Update().SetBic(splittedIDs[0]).
	// 			SetSerial(splittedIDs[1]).
	// 			SetChecksum(splittedIDs[2]).Save(ctx)
	// 	}

	// 	// tracking := &mygrpc.ContainerTracking{
	// 	// 	ID:          int32(s.ID),
	// 	// 	Score:       s.Score,
	// 	// 	ImageURL:    s.ImageURL,
	// 	// 	ContainerID: s.ContainerID,
	// 	// 	BIC:         splittedIDs[0],
	// 	// 	Serial:      splittedIDs[1],
	// 	// 	Checksum:    splittedIDs[2],
	// 	// 	CreatedAt:   int32(s.CreatedAt.Unix()),
	// 	// }

	// 	// resTrackings = append(resTrackings, tracking)
	// }

	// for rows.Next() {
	// 	t := &mygrpc.ContainerTracking{}
	// 	tzdate := &time.Time{}
	// 	err := rows.Scan(&t.ID, &t.ContainerID, &t.ImageURL, &t.Score, &t.BIC, &t.Serial, &t.Checksum, tzdate)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	t.CreatedAt = int32(tzdate.UTC().Unix())
	// 	resTrackings = append(resTrackings, t)
	// }

	for _, t := range tracks {
		tracking := &mygrpc.ContainerTracking{
			ID:          int32(t.ID),
			ContainerID: t.ContainerID,
			CreatedAt:   int32(t.CreatedAt.Unix()),
		}

		if len(t.Edges.Suggestions) > 0 {
			suggest := t.Edges.Suggestions[0]
			tracking.OCRID = int32(suggest.ID)
			tracking.ImageURL = suggest.ImageURL
			tracking.Score = suggest.Score
			tracking.BIC = suggest.Bic
			tracking.Serial = suggest.Serial
			tracking.Checksum = suggest.Checksum
		}

		resTrackings = append(resTrackings, tracking)
	}

	return &mygrpc.ResListContainerTrackings{
		Total:     int32(total),
		Trackings: resTrackings,
	}, nil
}

func (svc *Server) ListContainerOCRs(ctx context.Context, req *mygrpc.ReqEmpty) (*mygrpc.ResListContainerOCRs, error) {
	offset, limit := paging(req)

	total, err := svc.db.ContainerTrackingSuggestion.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	suggests, err := svc.db.ContainerTrackingSuggestion.Query().
		Order(ent.Desc(containertrackingsuggestion.FieldCreatedAt)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	resOCRs := []*mygrpc.ContainerOCR{}

	for _, s := range suggests {
		splittedIDs := splitContainerID(s.ContainerID)

		if s.Bic != splittedIDs[0] {
			s.Update().SetBic(splittedIDs[0]).
				SetSerial(splittedIDs[1]).
				SetChecksum(splittedIDs[2]).Save(ctx)
		}

		ocr := &mygrpc.ContainerOCR{
			ID:          int32(s.ID),
			Score:       s.Score,
			ImageURL:    s.ImageURL,
			ContainerID: s.ContainerID,
			BIC:         splittedIDs[0],
			Serial:      splittedIDs[1],
			Checksum:    splittedIDs[2],
			CreatedAt:   int32(s.CreatedAt.Unix()),
		}

		resOCRs = append(resOCRs, ocr)
	}

	return &mygrpc.ResListContainerOCRs{
		Total: int32(total),
		Ocrs:  resOCRs,
	}, nil
}

func (svc *Server) PullMLResult(req *mygrpc.ReqEmpty, resp mygrpc.MyGRPC_PullMLResultServer) error {
	// Create client UUID
	clientId := uuid.UUID(0)

	// Create client's data channel
	svc.clients[clientId] = make(chan *ContainerTrackingSuggestion, 10)
	defer func() {
		delete(svc.clients, clientId)
	}()

	// Send to browser
	for suggest := range svc.clients[clientId] {
		err := resp.Send(&mygrpc.ResMLResult{
			SuggestID:   suggest.ID,
			ContainerID: suggest.ContainerID,
			ImageURL:    suggest.ImageURL,
			Score:       suggest.Score,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *Server) ConfirmContainerID(ctx context.Context, req *mygrpc.ReqConfirmContainerID) (*mygrpc.ResEmpty, error) {
	builder := svc.db.ContainerTracking.Create().
		SetContainerID(req.ContainerID)

	if req.SuggestID != nil {
		builder = builder.AddSuggestionIDs(int(*req.SuggestID))
	}

	_, err := builder.Save(ctx)
	if err != nil {
		log.Println("!!!", err)
		return nil, err
	}

	return &mygrpc.ResEmpty{}, nil
}

var _ mygrpc.MyGRPCServer = &Server{}
