package entity

import (
	"errors"
	"reflect"
	"time"
)

type ProfilePic struct {
	Url  string `json:"url" firestore:"url"`
	Path string `json:"path" firestore:"path"`
}

type User struct {
	Id              string      `required:"true" json:"id" firestore:"id"`
	FirebaseAuthId  string      `json:"firebase_auth_id" firestore:"firebase_auth_id"`
	DisplayName     string      `required:"true" json:"display_name" firestore:"display_name"`
	AccountName     string      `required:"true" json:"account_name" firestore:"account_name"`
	FirstName       string      `json:"first_name" firestore:"first_name" firebase:"auth"`
	LastName        string      `json:"last_name" firestore:"last_name" firebase:"auth"`
	ProfilePic      *ProfilePic `json:"profile_picture" firestore:"profile_picture"`
	Age             int         `json:"age" firestore:"age" firebase:"auth"`
	Birthday        time.Time   `json:"birthday" firestore:"birthday" firebase:"auth"`
	Gender          string      `json:"gender" firestore:"gender" firebase:"auth"`
	SpouseId        string      `json:"spouse_id,omitempty" firestore:"spouse_id,omitempty" firebase:"auth"`
	ChildrenIds     []string    `json:"children_ids,omitempty" firestore:"children_ids,omitempty"`
	Address         *Address    `json:"address" firestore:"address" firebase:"auth"`
	HomeLocation    *Location   `json:"home_location" firestore:"home_location" firebase:"auth"`
	CurrentLocation *Location   `json:"current_location" firestore:"current_location"`
	CreatedAt       time.Time   `json:"created_at" firestore:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at" firestore:"updated_at"`
	FollowerIds     []string    `json:"follower_ids" firestore:"follower_ids"`
	FollowingIds    []string    `json:"following_ids" firestore:"following_ids"`
}

func (u *User) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

func (e *User) BeforeUpdate() {
	e.UpdatedAt = time.Now()
}

func (u *User) Validate() error {
	v := reflect.ValueOf(u).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("required")
		if tag == "true" && reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			return errors.New("required field is empty")
		}
	}
	return nil
}
