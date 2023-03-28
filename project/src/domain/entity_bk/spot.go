package entity

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Photo struct {
	Url       string `json:"url" firestore:"url"`
	OriginUrl string `json:"origin_url" firestore:"origin_url"`
	Path      string `json:"path" firestore:"path"`
	Provider  string `json:"provider" firestore:"provider"`
}

type Proprietary struct {
	Rating     int32    `json:"rating" firestore:"rating"`
	CheckInIds []string `json:"checkin_ids" firestore:"checkin_ids"`
}

type GoogleMap struct {
	PlaceId string    `json:"place_id" firestore:"place_id"`
	Types   []string  `json:"types" firestore:"types"`
	Reviews []*Review `json:"reviews" firestore:"reviews"`
	Rating  int32     `json:"rating" firestore:"rating"`
	Photos  []*Photo  `json:"photos" firestore:"photos"`
}

type Instagram struct {
	Types   []string  `json:"types" firestore:"types"`
	Reviews []*Review `json:"reviews" firestore:"reviews"`
	Rating  int32     `json:"rating" firestore:"rating"`
	Photos  []*Photo  `json:"photos" firestore:"photos"`
}

type Review struct {
	UserId    string    `json:"user_id,omitempty" firestore:"user_id,omitempty"`
	Rating    float64   `json:"rating" firestore:"rating"`
	Comment   string    `json:"comment" firestore:"comment"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at"`
}

func (r *Review) BeforeCreate() {
	now := time.Now()
	r.CreatedAt = now
	r.UpdatedAt = now
}

func (r *Review) BeforeUpdate() {
	r.UpdatedAt = time.Now()
}

type Spot struct {
	Id          string       `required:"true" json:"id" firestore:"id"`
	Name        string       `required:"true" json:"name" firestore:"name"`
	Description string       `json:"description" firestore:"description"`
	Provider    string       `required:"true" json:"provider" firestore:"provider"`
	Location    *Location    `required:"true" json:"location" firestore:"location"`
	Instagram   *Instagram   `json:"instagram,omitempty" firestore:"instagram,omitempty"`
	GoogleMap   *GoogleMap   `json:"google_map,omitempty" firestore:"google_map,omitempty"`
	Proprietary *Proprietary `json:"proprietary,omitempty" firestore:"proprietary,omitempty"`
	Address     *Address     `json:"address" firestore:"address"`
	CategoryIds []string     `json:"category_ids" firestore:"category_ids"`
	CreatedAt   time.Time    `required:"true" json:"created_at" firestore:"created_at"`
	UpdatedAt   time.Time    `required:"true" json:"updated_at" firestore:"updated_at"`
}

func (s *Spot) BeforeCreate() {
	now := time.Now()
	s.CreatedAt = now
	s.UpdatedAt = now
}

func (s *Spot) BeforeUpdate() {
	s.UpdatedAt = time.Now()
}

func (s *Spot) Validate() error {
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("required")
		if tag == "true" && reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			message := fmt.Sprintf("required field is empty: %v", v.Type().Field(i).Tag.Get("json"))
			return errors.New(message)
		}
	}
	return nil
}
