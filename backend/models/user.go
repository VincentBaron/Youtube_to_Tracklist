package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Friends  []User `gorm:"many2many:user_friends;"`
}
