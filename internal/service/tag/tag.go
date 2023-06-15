package tag

import (
	"github.com/mephistolie/chefbook-backend-tag/internal/entity"
)

func (s *Service) GetTagsAndGroups(languageCode string, groupIds *[]string) ([]entity.Tag, map[string]string) {
	return s.repo.GetTagsAndGroups(languageCode, groupIds)
}

func (s *Service) GetTagsMapWithGroups(tagIds []string, languageCode string) (map[string]entity.Tag, map[string]string) {
	return s.repo.GetTagsMapWithGroups(tagIds, languageCode)
}

func (s *Service) GetTagWithGroup(tagId, languageCode string) (entity.Tag, *string, error) {
	return s.repo.GetTagWithGroup(tagId, languageCode)
}

func (s *Service) GetGroups(languageCode string) map[string]string {
	return s.repo.GetGroups(languageCode)
}
