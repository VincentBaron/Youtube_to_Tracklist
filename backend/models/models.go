// Definition of the structures and SQL interaction functions
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	SpotifyID string
	Playlists []Playlist
}

type Playlist struct {
	gorm.Model
	SpotifyID string
	UserID    uint
	Name      string
	Tracks    []Track
}

type Track struct {
	gorm.Model
	SpotifyID  string
	PlaylistID uint
	Name       string
	Link       string
}

// Models return one of every model the database must create..
func Models() []interface{} {
	return []interface{}{
		&Playlist{},
		&Track{},
	}
}
