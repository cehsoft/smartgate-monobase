package server

import (
	"context"
	"log"

	"github.com/leandro-lugaresi/hub"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/model"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

type Server struct {
	hub             *hub.Hub
	db              *ent.Client
	cachedTrackings []*model.ContainerTracking

	mygrpc.UnimplementedMyGRPCServer
}

func CreateServer(db *ent.Client) *Server {
	hub := hub.New()

	return &Server{
		hub:             hub,
		db:              db,
		cachedTrackings: []*model.ContainerTracking{},
	}
}

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req)

	svc.hub.Publish(hub.Message{
		Name: "ml.new.result",
		Fields: hub.Fields{
			"CachedID":    len(svc.cachedTrackings),
			"ContainerID": req.ContainerID,
			"ImageURL":    req.ImageURL,
			"Score":       req.Score,
		},
	})

	svc.cachedTrackings = append(svc.cachedTrackings, &model.ContainerTracking{
		ContainerID: req.ContainerID,
		ImageURL:    req.ImageURL,
	})

	return &mygrpc.ResEmpty{}, nil
}

func (svc *Server) PullMLResult(req *mygrpc.ReqEmpty, resp mygrpc.MyGRPC_PullMLResultServer) error {
	sub := svc.hub.Subscribe(100, "ml.new.result")

	for msg := range sub.Receiver {
		var ContainerID, ImageURL string
		var Score float32
		var CachedID int
		var ok bool

		if ContainerID, ok = msg.Fields["ContainerID"].(string); !ok {
			continue
		}

		if ImageURL, ok = msg.Fields["ImageURL"].(string); !ok {
			continue
		}

		if Score, ok = msg.Fields["Score"].(float32); !ok {
			continue
		}

		if CachedID, ok = msg.Fields["CachedID"].(int); !ok {
			continue
		}

		err := resp.Send(&mygrpc.ResMLResult{
			CachedID:    int32(CachedID),
			ContainerID: ContainerID,
			ImageURL:    ImageURL,
			Score:       Score,
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

	if req.CachedID != nil && len(svc.cachedTrackings) > int(*req.CachedID) {
		cached := svc.cachedTrackings[*req.CachedID]
		builder = builder.
			SetImageURL(cached.ImageURL)
		if cached.ContainerID != req.ContainerID {
			builder = builder.
				SetManual(true).
				SetImageURL(cached.ImageURL)
		}
	}

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	svc.cachedTrackings = []*model.ContainerTracking{}

	return &mygrpc.ResEmpty{}, nil
}

var _ mygrpc.MyGRPCServer = &Server{}
