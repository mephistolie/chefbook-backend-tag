package dto

import (
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-tag/internal/entity"
)

func NewGetTagsResponse(tags []entity.Tag, groups map[string]string) *api.GetTagsResponse {
	dtos := make([]*api.Tag, len(tags))
	for i, tag := range tags {
		dtos[i] = NewTag(tag)
	}
	return &api.GetTagsResponse{Tags: dtos, GroupNames: groups}
}

func NewGetTagsMapResponse(tags map[string]entity.Tag, groups map[string]string) *api.GetTagsMapResponse {
	dtos := make(map[string]*api.Tag)
	for id, tag := range tags {
		dtos[id] = NewTag(tag)
	}
	return &api.GetTagsMapResponse{Tags: dtos, GroupNames: groups}
}

func NewTag(tag entity.Tag) *api.Tag {
	return &api.Tag{
		TagId:   tag.Id,
		Name:    *tag.Name,
		Emoji:   tag.Emoji,
		GroupId: tag.Emoji,
	}
}
