package controller

import (
	"app/usecase/port"
	"context"
)

type SpotController struct {
	spotInputPort port.SpotInputPort
}

func NewSpotContoller(fi port.SpotInputPort) *SpotController {
	return &SpotController{
		spotInputPort: fi,
	}
}

func (fc *SpotController) GetSpotDataFromGoogleMap(ctx context.Context, searchText string) {
	fc.spotInputPort.GetSpotDataFromGoogleMap(ctx, searchText)
}

func (fc *SpotController) GetAllSpots(ctx context.Context) (*port.BaseJsonMultipleResponse, error) {
	return fc.spotInputPort.GetAllSpots(ctx)
}
