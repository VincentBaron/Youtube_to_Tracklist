package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Title  string  `json:"title,omitempty"`
	Tracks []Track `gorm:"many2many:playlist_tracks;"`
}
