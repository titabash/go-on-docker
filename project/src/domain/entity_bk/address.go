package entity

type Address struct {
	Line1       string `json:"line1" firestore:"line1"`
	Line2       string `json:"line2" firestore:"line2"`
	City        string `json:"city" firestore:"city"`
	State       string `json:"state" firestore:"state"`
	PostalCode  string `json:"postal_code" firestore:"postal_code"`
	Country     string `json:"country" firestore:"country"`
	CountryCode string `json:"country_code" firestore:"country_code"`
}
