package server

import (
	"context"
	"log"

	"github.com/leandro-lugaresi/hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/model"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

type Server struct {
	hub             *hub.Hub
	cachedTrackings []*model.ContainerTracking

	mygrpc.UnimplementedMyGRPCServer
}

func CreateServer() *Server {
	hub := hub.New()

	return &Server{
		hub:             hub,
		cachedTrackings: []*model.ContainerTracking{},
	}
}

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req.ContainerID)

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
	sub := svc.hub.Subscribe(10, "ml.new.result")

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

		resp.Send(&mygrpc.ResMLResult{
			CachedID:    int32(CachedID),
			ContainerID: ContainerID,
			ImageURL:    ImageURL,
			Score:       Score,
		})
	}

	return nil
}

func (svc *Server) ConfirmContainerID(context.Context, *mygrpc.ReqConfirmContainerID) (*mygrpc.ResEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmContainerID not implemented")
}

var _ mygrpc.MyGRPCServer = &Server{}
