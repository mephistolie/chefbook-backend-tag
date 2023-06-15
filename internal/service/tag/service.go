package tag

import (
	"github.com/mephistolie/chefbook-backend-tag/internal/service/dependencies/repository"
)

type Service struct {
	repo repository.Tag
}

func NewService(
	repo repository.Tag,
) *Service {
	return &Service{repo: repo}
}
