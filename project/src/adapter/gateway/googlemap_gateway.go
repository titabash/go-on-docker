package gateway

import (
	"app/infra/googlemap"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type GoogleMapGateway struct {
	placeDetailsApi *googlemap.GooglePlaceDetailsApi
	placeSearchApi  *googlemap.GooglePlaceTextSearchApi
}

func NewGoogleMapGateway() *GoogleMapGateway {
	return &GoogleMapGateway{
		placeDetailsApi: googlemap.NewGooglePlaceDetailsApi(),
		placeSearchApi:  googlemap.NewGooglePlaceTextSearchApi(),
	}
}

func (gg *GoogleMapGateway) GetPlaceDetailsByPlaceId(ctx context.Context, placeId string) (*googlemap.PlaceDetailsResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet, gg.placeDetailsApi.Endpoint, nil)

	q := req.URL.Query()
	q.Set("place_id", placeId)
	q.Set("region", "jp")
	q.Set("language", "ja")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, err
	}
	resp, err := gg.placeDetailsApi.Client.Do(req)

	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)

	defer resp.Body.Close()
	var response *googlemap.PlaceDetailsResponse
	json.Unmarshal(body, &response)

	return response, nil
}

func (gg *GoogleMapGateway) GetPlacesByText(ctx context.Context, pageToken string, searchText string) (*googlemap.TextSearchResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet, gg.placeSearchApi.Endpoint, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if pageToken == "" {
		q.Set("query", searchText)
	} else {
		q.Set("pagetoken", pageToken)
	}
	q.Set("region", "jp")
	q.Set("language", "ja")
	req.URL.RawQuery = q.Encode()

	resp, err := gg.placeSearchApi.Client.Do(req)

	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var response *googlemap.TextSearchResponse
	json.Unmarshal(body, &response)

	return response, nil
}
