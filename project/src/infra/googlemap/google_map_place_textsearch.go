package googlemap

import (
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
)

type TextSearchResult struct {
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Viewport struct {
			Northeast struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"northeast"`
			Southwest struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"viewport"`
	} `json:"geometry"`
	Icon                string   `json:"icon"`
	IconBackgroundColor string   `json:"icon_background_color"`
	IconMaskBaseURI     string   `json:"icon_mask_base_uri"`
	Name                string   `json:"name"`
	PlaceID             string   `json:"place_id"`
	Reference           string   `json:"reference"`
	Types               []string `json:"types"`
}

type TextSearchResponse struct {
	HTMLAttributions []interface{}      `json:"html_attributions"`
	NextPageToken    string             `json:"next_page_token"`
	Results          []TextSearchResult `json:"results"`
	Status           string             `json:"status"`
}

type GooglePlaceTextSearchApi struct {
	Client   *http.Client
	Endpoint string
}

func NewGooglePlaceTextSearchApi() *GooglePlaceTextSearchApi {
	return &GooglePlaceTextSearchApi{
		Endpoint: "https://maps.googleapis.com/maps/api/place/textsearch/json",
		Client: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: &transport.APIKey{Key: os.Getenv("GOOGLE_MAPS_API_KEY")},
		},
	}
}
