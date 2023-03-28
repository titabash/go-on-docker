package entity

type Location struct {
	Lat float64 `required:"true" json:"lat" firestore:"lat"`
	Lng float64 `required:"true" json:"lng" firestore:"lng"`
}
