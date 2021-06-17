package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

// Server is the Logic handler for the server
// It has to fullfill the GRPC schema generated Interface
// In this case its only 1 function called Ping
type Server struct {
	mygrpc.UnimplementedMyGRPCServer
}

func (svc *Server) NewMLResult(ctx context.Context, req *mygrpc.ReqMLResult) (*mygrpc.ResEmpty, error) {
	log.Println("ML result received", req.ContainerID)
	return &mygrpc.ResEmpty{}, nil
}

func (svc *Server) PullMLResult(req *mygrpc.ReqEmpty, resp mygrpc.MyGRPC_PullMLResultServer) error {
	cnt := 0

	for {
		cnt += 1

		resp.Send(&mygrpc.ResMLResult{
			ContainerID: fmt.Sprintf("MALU-XXYY-%d", cnt),
		})

		time.Sleep(5 * time.Second)
	}

}

func (svc *Server) ConfirmContainerID(context.Context, *mygrpc.ReqConfirmContainerID) (*mygrpc.ResConfirmContainerID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmContainerID not implemented")
}

var _ mygrpc.MyGRPCServer = &Server{}
