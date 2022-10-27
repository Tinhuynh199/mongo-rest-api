package model

import (
	"time"
)

type Student struct {
	ID          string     `json:"id" bson:"_id"`
	Username    string     `json:"username" bson:"username"`
	Email       string     `json:"email" bson:"email"`
	Phone       string     `json:"phone" bson:"phone"`
	DateOfBirth *time.Time `json:"dateOfBirth" bson:"dateOfBirth"`
	Course      string     `json:"course" bson:"course"`
}
