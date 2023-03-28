package googlemap

import (
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type Geometry struct {
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Photo struct {
	// Height           int      `json:"height"`
	// HtmlAttributions []string `json:"html_attributions"`
	PhotoReference string `json:"photo_reference"`
	// Width            int      `json:"width"`
}

type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Result struct {
	AddressComponents []AddressComponents `json:"address_components"`
	AdrAddress        string              `json:"adr_address"`
	BusinessStatus    string              `json:"business_status"`
	FormattedAddress  string              `json:"formatted_address"`
	Geometry          Geometry            `json:"geometry"`
	Icon              string              `json:"icon"`
	Name              string              `json:"name"`
	OpeningHours      OpeningHours        `json:"opening_hours"`
	Photos            []Photo             `json:"photos"`
	PlaceID           string              `json:"place_id"`
	PlusCode          struct {
		CompoundCode string `json:"compound_code"`
		GlobalCode   string `json:"global_code"`
	} `json:"plus_code"`
	Rating      float32  `json:"rating"`
	Reference   string   `json:"reference"`
	Reviews     []Review `json:"reviews"`
	Scope       string   `json:"scope"`
	Types       []string `json:"types"`
	Url         string   `json:"url"`
	UserRatings int      `json:"user_ratings_total"`
	Vicinity    string   `json:"vicinity"`
	Website     string   `json:"website"`
}

type Review struct {
	AuthorName      string  `json:"author_name"`
	AuthorURL       string  `json:"author_url"`
	Language        string  `json:"language"`
	ProfilePhotoURL string  `json:"profile_photo_url"`
	Rating          float64 `json:"rating"`
	RelativeTime    string  `json:"relative_time_description"`
	Text            string  `json:"text"`
	Time            int     `json:"time"`
}

type PlaceDetailsResponse struct {
	Result Result `json:"result"`
	Status string `json:"status"`
}

type GooglePlaceDetailsApi struct {
	Client   *http.Client
	Endpoint string
}

func NewGooglePlaceDetailsApi() *GooglePlaceDetailsApi {
	return &GooglePlaceDetailsApi{
		Endpoint: "https://maps.googleapis.com/maps/api/place/details/json",
		Client: &http.Client{
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
			Transport: &transport.APIKey{Key: os.Getenv("GOOGLE_MAPS_API_KEY")},
		},
	}
}
