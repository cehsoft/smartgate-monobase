package server

import (
	"context"
	"log"
	"time"

	"github.com/init-tech-solution/service-spitc-stream/services/lib/uuid"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/model"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

type Server struct {
	db *ent.Client

	clients map[string]chan *ContainerTrackingSuggestion

	mygrpc.UnimplementedMyGRPCServer
}

type ContainerTrackingSuggestion struct {
	ID int32
	model.ContainerTracking
}

func CreateServer(db *ent.Client) *Server {
	return &Server{
		db:      db,
		clients: map[string]chan *ContainerTrackingSuggestion{},
	}
}

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req)

	suggestRrd, err := svc.db.ContainerTrackingSuggestion.Create().
		SetContainerID(req.ContainerID).
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
	tracks, err := svc.db.ContainerTracking.Query().WithSuggestions().
		Order(ent.Desc(containertracking.FieldCreatedAt)).
		Limit(100).
		All(ctx)
	if err != nil {
		return nil, err
	}

	resTrackings := []*mygrpc.ContainerTracking{}

	for _, t := range tracks {
		tracking := &mygrpc.ContainerTracking{
			ID:          int32(t.ID),
			ContainerID: t.ContainerID,
			CreatedAt:   int32(t.CreatedAt.Unix()),
		}

		if len(t.Edges.Suggestions) > 0 {
			suggest := t.Edges.Suggestions[0]
			tracking.ImageURL = suggest.ImageURL
		}

		resTrackings = append(resTrackings, tracking)
	}

	return &mygrpc.ResListContainerTrackings{
		Trackings: resTrackings,
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
