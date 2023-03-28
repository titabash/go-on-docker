package repository

import (
	"app/domain/entity"
	"context"
)

type SpotRepository interface {
	NewSpotId() string
	FindByDocumentID(ctx context.Context, id string) (*entity.Spot, error)
	FindBetweenLatLng(ctx context.Context, northEast *entity.Location, southWest *entity.Location) ([]*entity.Spot, error)
	UploadImage(ctx context.Context, file []byte, provider string, spotId string, name string) (string, string, error)
	FindAll(ctx context.Context, lastSpotId string) ([]*entity.Spot, error)
	FindByGooglePlaceID(ctx context.Context, googlePlaceId string) ([]*entity.Spot, error)
	InsertOne(ctx context.Context, spot *entity.Spot) error
	Update(ctx context.Context, id string, spot *entity.Spot) error
	Closer() error
}
