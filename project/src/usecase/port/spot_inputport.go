package port

import "context"

type SpotInputPort interface {
	GetSpotDataFromGoogleMap(ctx context.Context, searchText string) error
	GetAllSpots(ctx context.Context) (*BaseJsonMultipleResponse, error)
}
