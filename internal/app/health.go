package app

import (
	"context"

	"github.com/jmoiron/sqlx"
	tagpb "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	eventlog "github.com/mephistolie/chefbook-backend-tag/internal/logging"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"time"
)

func monitorHealthChecking(ctx context.Context, db *sqlx.DB, healthServer *health.Server) {
	events := eventlog.NewEvents()
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		status := healthpb.HealthCheckResponse_SERVING
		if err := db.PingContext(ctx); err != nil {
			status = healthpb.HealthCheckResponse_NOT_SERVING
			events.PostgresHealthCheckFailed(ctx, err)
		}
		setHealthStatus(healthServer, status)

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func setHealthStatus(healthServer *health.Server, status healthpb.HealthCheckResponse_ServingStatus) {
	healthServer.SetServingStatus("", status)
	healthServer.SetServingStatus(tagpb.TagService_ServiceDesc.ServiceName, status)
}
