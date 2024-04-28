package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Title  string `json:"title,omitempty"`
	Artist string `json:"artist,omitempty"`
	Link   string `json:"link,omitempty"`
}
