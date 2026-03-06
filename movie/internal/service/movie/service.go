package movie

import (
	"context"
	"errors"
	metadataModel "movieapp/metadata/pkg"
	"movieapp/movie/internal/gateway"
	model "movieapp/movie/pkg"
	ratingModel "movieapp/rating/pkg"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingModel.RecordID, recordType ratingModel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingModel.RecordID, recordType ratingModel.RecordType, rating *ratingModel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadataModel.Metadata, error)
}

type Service struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, memetadataGateway metadataGateway) *Service {
	return &Service{ratingGateway, memetadataGateway}
}

func (s *Service) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := s.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}

	rating, err := s.ratingGateway.GetAggregatedRating(ctx, ratingModel.RecordID(id), ratingModel.RecordTypeMovie)

	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		return nil, err
	}

	details.Rating = &rating

	return details, nil
}
