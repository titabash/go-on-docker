package entity

import (
	"errors"
	"reflect"
	"time"
)

type CheckIn struct {
	Id        string    `required:"true" json:"id" firestore:"id"`
	UserId    string    `required:"true" json:"user_id" firestore:"user_id"`
	SpotId    string    `required:"true" json:"spot_id" firestore:"spot_id"`
	Photos    []*Photo  `required:"true" json:"photos" firestore:"photos"`
	Address   *Address  `required:"true" json:"address" firestore:"address"`
	Location  *Location `required:"true" json:"location" firestore:"location"`
	CreatedAt time.Time `required:"true" json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time `required:"true" json:"updated_at" firestore:"updated_at"`
}

func (c *CheckIn) BeforeCreate() {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
}

func (c *CheckIn) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}

func (c *CheckIn) Validate() error {
	v := reflect.ValueOf(c).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("required")
		if tag == "true" && reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			return errors.New("required field is empty")
		}
	}
	return nil
}
