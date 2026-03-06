package metadata

import (
	"context"
	"errors"
	"movieapp/metadata/internal/repository"
	model "movieapp/metadata/pkg"
)

// ErrNotFound is returned when a requested record is not found
var ErrNotFound = errors.New("not found")

// metadataRepository is a wrapper around the repository.
type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
	Put(ctx context.Context, metadata *model.Metadata) error
}

type Service struct {
	repo metadataRepository
}

func New(repo metadataRepository) *Service {
	return &Service{repo}
}

// Get returns movie metadata by id
func (s *Service) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := s.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}

func (s *Service) Put(ctx context.Context, metadata *model.Metadata) error {
	return s.repo.Put(ctx, metadata)
}
