package main

import (
	"log"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/server"
)

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	apiserver := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	lis, err := net.Listen("tcp", "127.0.0.1:9990")
	if err != nil {
		log.Fatal(err)
	}

	server := server.CreateServer()
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
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	sv.Start(":3000")
}
