package logging

import (
	"context"

	"github.com/mephistolie/chefbook-backend-common/log"
)

type Events struct{}

func NewEvents() Events {
	return Events{}
}

func (Events) ConfigLoaded(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "config.loaded",
		Message:   "service configuration loaded",
		Component: "config",
	})
}

type StartupFailure struct {
	Operation string
}

func (Events) StartupFailed(ctx context.Context, data StartupFailure, err error) {
	log.LogFatal(ctx, log.Event{
		Event:     "app.startup.failed",
		Message:   "service startup failed",
		Component: "app",
		Operation: data.Operation,
	}, err)
}

func (Events) GRPCServerStarted(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "grpc.server.started",
		Message:   "gRPC server started",
		Component: log.ComponentGRPC,
	})
}

func (Events) GRPCServerFailed(ctx context.Context, err error) {
	log.LogError(ctx, log.Event{
		Event:     "grpc.server.failed",
		Message:   "gRPC server stopped with an error",
		Component: log.ComponentGRPC,
	}, err)
}

func (Events) PostgresHealthCheckFailed(ctx context.Context, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "postgres.health_check.failed",
		Message:   "database is unavailable",
		Component: log.ComponentPostgres,
	}, err)
}

type PostgresOperation struct {
	Operation string
	Entity    string
}

func (Events) PostgresQueryFailed(ctx context.Context, data PostgresOperation, err error) {
	log.LogError(ctx, log.Event{
		Event:     "postgres.query.failed",
		Message:   "database query failed",
		Component: log.ComponentPostgres,
		Operation: data.Operation,
		Payload: map[string]any{
			"entity": data.Entity,
		},
	}, err)
}

func (Events) PostgresRowScanFailed(ctx context.Context, data PostgresOperation, err error) {
	log.LogError(ctx, log.Event{
		Event:     "postgres.row.scan_failed",
		Message:   "database row scan failed",
		Component: log.ComponentPostgres,
		Operation: data.Operation,
		Payload: map[string]any{
			"entity": data.Entity,
		},
	}, err)
}

func (Events) PostgresRowsIterationFailed(ctx context.Context, data PostgresOperation, err error) {
	log.LogError(ctx, log.Event{
		Event:     "postgres.rows.iteration_failed",
		Message:   "database rows iteration failed",
		Component: log.ComponentPostgres,
		Operation: data.Operation,
		Payload: map[string]any{
			"entity": data.Entity,
		},
	}, err)
}

type TagLookup struct {
	TagID        string
	LanguageCode string
}

func (Events) TagLookupFailed(ctx context.Context, data TagLookup, err error) {
	log.LogError(ctx, log.Event{
		Event:     "tag.lookup.failed",
		Message:   "tag lookup failed",
		Component: log.ComponentPostgres,
		Operation: "get_tag",
		Payload: map[string]any{
			"tag_id": data.TagID,
		},
	}, err)
}

func (Events) TagNotFound(ctx context.Context, data TagLookup) {
	log.LogDebug(ctx, log.Event{
		Event:     "tag.not_found",
		Message:   "tag not found",
		Component: log.ComponentPostgres,
		Operation: "get_tag",
		Payload: map[string]any{
			"tag_id": data.TagID,
		},
	})
}

func (Events) TagNameUnavailable(ctx context.Context, data TagLookup) {
	log.LogWarn(ctx, log.Event{
		Event:     "tag.name.unavailable",
		Message:   "tag name is unavailable for requested language",
		Component: log.ComponentPostgres,
		Operation: "get_tag",
		Payload: map[string]any{
			"tag_id":        data.TagID,
			"language_code": data.LanguageCode,
		},
	})
}
