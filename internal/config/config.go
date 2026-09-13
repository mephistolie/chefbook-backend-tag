package config

import (
	"context"

	"github.com/mephistolie/chefbook-backend-tag/internal/logging"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	Database Database
}

type Database struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	DBName   *string
}

func (c Config) Validate() error {
	if *c.Environment != EnvProd {
		*c.Environment = EnvDev
	}
	return nil
}

func (c Config) Print(ctx context.Context) {
	logging.NewEvents().ConfigLoaded(ctx)
}
