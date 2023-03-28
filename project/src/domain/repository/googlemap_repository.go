package repository

import (
	"app/infra/googlemap"
	"context"
)

type GoogleMapRepository interface {
	GetPlaceDetailsByPlaceId(ctx context.Context, placeId string) (*googlemap.PlaceDetailsResponse, error)
	GetPlacesByText(ctx context.Context, pageToken string, searchText string) (*googlemap.TextSearchResponse, error)
}
