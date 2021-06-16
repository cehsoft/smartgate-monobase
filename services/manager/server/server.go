package server

import (
	"github.com/init-tech-solution/service-spitc-stream/mygrpc"
)

// Server is the Logic handler for the server
// It has to fullfill the GRPC schema generated Interface
// In this case its only 1 function called Ping
type Server struct {
	mygrpc.UnimplementedMyGRPCServer
}
