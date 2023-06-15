package app

import (
	"context"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	tagpb "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-tag/internal/config"
	"github.com/mephistolie/chefbook-backend-tag/internal/repository/postgres"
	service "github.com/mephistolie/chefbook-backend-tag/internal/service/tag"
	tag "github.com/mephistolie/chefbook-backend-tag/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"
)

func Run(cfg *config.Config) {
	log.Init(*cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	db, err := postgres.Connect(cfg.Database)
	if err != nil {
		log.Fatal(err)
		return
	}

	repository := postgres.NewRepository(db)

	tagService := service.NewService(repository)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			log.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	tagServer := tag.NewServer(tagService)

	go monitorHealthChecking(db, healthServer)

	tagpb.RegisterTagServiceServer(grpcServer, tagServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		} else {
			log.Info("gRPC server started")
		}
	}()

	wait := shutdown.Graceful(context.Background(), 5*time.Second, map[string]shutdown.Operation{
		"grpc-server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			return db.Close()
		},
	})
	<-wait
}
