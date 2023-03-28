package entity

type Category struct {
	Id             string   `required:"true" json:"id" firestore:"id"`
	Name           string   `required:"true" json:"name" firestore:"name"`
	GoogleMapTypes []string `json:"google_map_type" firestore:"google_map_type"`
}
