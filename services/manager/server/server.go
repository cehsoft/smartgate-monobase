package server

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/model"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

type Server struct {
	db *ent.Client

	cachedTrackings []*CachedContainerTracking
	clients         map[string]chan *CachedContainerTracking

	mygrpc.UnimplementedMyGRPCServer
}

type CachedContainerTracking struct {
	CacheID int
	model.ContainerTracking
}

func CreateServer(db *ent.Client) *Server {
	return &Server{
		db:              db,
		clients:         map[string]chan *CachedContainerTracking{},
		cachedTrackings: []*CachedContainerTracking{},
	}
}

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req)

	cached := CachedContainerTracking{
		CacheID: len(svc.cachedTrackings),
		ContainerTracking: model.ContainerTracking{
			ContainerID: req.ContainerID,
			ImageURL:    req.ImageURL,
			Score:       req.Score,
		},
	}

	// Append to cached
	svc.cachedTrackings = append(svc.cachedTrackings, &cached)

	// Send to subscribed clients
	for uuid, trackingC := range svc.clients {
		go func(uuid string, trackingC chan *CachedContainerTracking) {
			timeout := time.NewTimer(10 * time.Second)

			select {
			case trackingC <- &cached:
				log.Println(">>>> sent:", uuid)
			case <-timeout.C:
				log.Println("!!! timeout:", uuid)
			}
		}(uuid, trackingC)
	}

	return &mygrpc.ResEmpty{}, nil
}

func (svc *Server) PullMLResult(req *mygrpc.ReqEmpty, resp mygrpc.MyGRPC_PullMLResultServer) error {
	// Create client UUID
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	// Create client's data channel
	svc.clients[uuid] = make(chan *CachedContainerTracking, 10)
	defer func() {
		delete(svc.clients, uuid)
	}()

	// Send to browser
	for ocr := range svc.clients[uuid] {
		err := resp.Send(&mygrpc.ResMLResult{
			CachedID:    int32(ocr.CacheID),
			ContainerID: ocr.ContainerID,
			ImageURL:    ocr.ImageURL,
			Score:       ocr.Score,
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
				SetManual(true)
		}
	}

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	// TODO better clear cache logics
	svc.cachedTrackings = []*CachedContainerTracking{}

	return &mygrpc.ResEmpty{}, nil
}

var _ mygrpc.MyGRPCServer = &Server{}
