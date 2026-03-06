package memory

import (
	"context"
	"movieapp/metadata/internal/repository"
	model "movieapp/metadata/pkg"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrieves movies metadata by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

// Put adds movies metadata for a given movie id.
func (r *Repository) Put(_ context.Context, metadata *model.Metadata) error {
	r.RLock()
	defer r.RUnlock()

	r.data[metadata.ID] = metadata

	return nil
}
