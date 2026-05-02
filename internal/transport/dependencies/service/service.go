package service

import (
	"context"

	"github.com/mephistolie/chefbook-backend-tag/internal/entity"
)

type Service interface {
	GetTagsAndGroups(ctx context.Context, languageCode string, groupIds *[]string) ([]entity.Tag, map[string]string)
	GetTagsMapWithGroups(ctx context.Context, tagIds []string, languageCode string) (map[string]entity.Tag, map[string]string)
	GetTagWithGroup(ctx context.Context, tagId, languageCode string) (entity.Tag, *string, error)
	GetGroups(ctx context.Context, languageCode string) map[string]string
}
