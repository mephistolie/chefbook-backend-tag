package tag

import (
	"context"

	"github.com/mephistolie/chefbook-backend-tag/internal/entity"
)

func (s *Service) GetTagsAndGroups(ctx context.Context, languageCode string, groupIds *[]string) ([]entity.Tag, map[string]string) {
	return s.repo.GetTagsAndGroups(ctx, languageCode, groupIds)
}

func (s *Service) GetTagsMapWithGroups(ctx context.Context, tagIds []string, languageCode string) (map[string]entity.Tag, map[string]string) {
	return s.repo.GetTagsMapWithGroups(ctx, tagIds, languageCode)
}

func (s *Service) GetTagWithGroup(ctx context.Context, tagId, languageCode string) (entity.Tag, *string, error) {
	return s.repo.GetTagWithGroup(ctx, tagId, languageCode)
}

func (s *Service) GetGroups(ctx context.Context, languageCode string) map[string]string {
	return s.repo.GetGroups(ctx, languageCode)
}
