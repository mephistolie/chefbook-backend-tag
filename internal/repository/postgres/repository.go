package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mephistolie/chefbook-backend-tag/internal/config"
	"k8s.io/utils/strings/slices"
)

const (
	tagsTable   = "tags"
	groupsTable = "tag_groups"
)

type Repository struct {
	db *sqlx.DB
}

func Connect(cfg config.Database) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=require",
			*cfg.Host, *cfg.Port, *cfg.User, *cfg.DBName, *cfg.Password))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getNameColumn(languageCode string) string {
	if slices.Contains(supportedLanguages, languageCode) {
		return r.formatNameColumn(languageCode)
	}
	return r.getFallbackNameColumn(languageCode)
}

func (r *Repository) getFallbackNameColumn(languageCode string) string {
	if slices.Contains(ruConsonantLanguages, languageCode) {
		return r.formatNameColumn(codeRu)
	}

	return r.formatNameColumn(codeEn)
}

func (r *Repository) formatNameColumn(languageCode string) string {
	return fmt.Sprintf("%s_%s", nameColumnPrefix, languageCode)
}
