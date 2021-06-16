package main

import (
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/server"
)

// GenerateTLSApi will load TLS certificates and key and create a grpc server with those.
func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(
		grpc.Creds(cred),
	)
	return s, nil
}

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	log.Println(">>>> here")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}

		if m.IsGrpcWebSocketRequest(r) {
			m.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	apiserver, err := GenerateTLSApi("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	// Start listening on a TCP Port
	lis, err := net.Listen("tcp", "127.0.0.1:9990")
	if err != nil {
		log.Fatal(err)
	}

	server := &server.Server{}
	mygrpc.RegisterMyGRPCServer(apiserver, server)

	go func() {
		log.Fatal(apiserver.Serve(lis))
	}()

	grpcWebServer := grpcweb.WrapServer(apiserver, grpcweb.WithWebsockets(true))
	multiplex := grpcMultiplexer{
		grpcWebServer,
	}

	sv := echo.New()
	sv.Use(middleware.Logger())
	sv.Use(middleware.Recover())
	sv.Use(middleware.CORS())
	sv.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if multiplex.IsGrpcWebSocketRequest(c.Request()) {
				multiplex.HandleGrpcWebsocketRequest(c.Response(), c.Request())
				return nil
			}

			return next(c)
		}
	})

	sv.GET("*", func(c echo.Context) error {
		return nil
	})

	sv.StartTLS(":3002", "cert/server.crt", "cert/server.key")
}
