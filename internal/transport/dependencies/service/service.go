package service

import (
	"github.com/mephistolie/chefbook-backend-tag/internal/entity"
)

type Service interface {
	GetTagsAndGroups(languageCode string, groupIds *[]string) ([]entity.Tag, map[string]string)
	GetTagsMapWithGroups(tagIds []string, languageCode string) (map[string]entity.Tag, map[string]string)
	GetTagWithGroup(tagId, languageCode string) (entity.Tag, *string, error)
	GetGroups(languageCode string) map[string]string
}
