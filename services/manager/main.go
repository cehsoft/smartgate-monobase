package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"context"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/migrate"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/server"
)

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

type Config struct {
	DB       string `env:"DB,required"`
	TLS      bool   `env:"TLS,default=false"`
	HTTP     string `env:"HTTP_PORT,default=:3000"`
	GRPC     string `env:"GRPC_PORT,default=:9990"`
	CertPath string `env:"CERT_PATH,default=cert/"`
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
	ctx := context.Background()

	var config Config
	l := envconfig.PrefixLookuper("SPITC_MANAGER_", envconfig.OsLookuper())
	if err := envconfig.ProcessWith(ctx, &config, l); err != nil {
		log.Fatalln(err)
	}
	cf, _ := json.Marshal(config)
	log.Println("Loaded config:", string(cf))

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

	lis, err := net.Listen("tcp", config.GRPC)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("pgx", config.DB)
	if err != nil {
		log.Fatal(err)
	}

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db.DB)))
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		// migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	server := server.CreateServer(client, db)
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

			if multiplex.IsGrpcWebRequest(c.Request()) {
				multiplex.HandleGrpcWebRequest(c.Response(), c.Request())
				return nil
			}

			return next(c)
		}
	})

	sv.GET("*", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	if config.TLS {
		err = sv.StartTLS(config.HTTP, config.CertPath+"server.crt", config.CertPath+"server.key")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err = sv.Start(config.HTTP)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
