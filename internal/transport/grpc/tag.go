package grpc

import (
	"context"
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-tag/internal/transport/grpc/dto"
)

func (s *TagServer) GetTags(_ context.Context, req *api.GetTagsRequest) (*api.GetTagsResponse, error) {
	var groups *[]string
	if len(req.Groups) > 0 {
		groups = &req.Groups
	}
	tags, groupNames := s.service.GetTagsAndGroups(req.LanguageCode, groups)

	return dto.NewGetTagsResponse(tags, groupNames), nil
}

func (s *TagServer) GetTagsMap(_ context.Context, req *api.GetTagsMapRequest) (*api.GetTagsMapResponse, error) {
	tags, groupNames := s.service.GetTagsMapWithGroups(req.TagIds, req.LanguageCode)
	return dto.NewGetTagsMapResponse(tags, groupNames), nil
}

func (s *TagServer) GetTag(_ context.Context, req *api.GetTagRequest) (*api.GetTagResponse, error) {
	tag, groupNamePtr, err := s.service.GetTagWithGroup(req.TagId, req.LanguageCode)
	if err != nil {
		return nil, err
	}
	groupName := ""
	if groupNamePtr != nil {
		groupName = *groupNamePtr
	}

	return &api.GetTagResponse{Tag: dto.NewTag(tag), GroupName: groupName}, nil
}

func (s *TagServer) GetTagGroups(_ context.Context, req *api.GetTagGroupsRequest) (*api.GetTagGroupsResponse, error) {
	groups := s.service.GetGroups(req.LanguageCode)
	return &api.GetTagGroupsResponse{Groups: groups}, nil
}
