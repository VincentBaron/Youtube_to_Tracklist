// Definition of the structures and SQL interaction functions
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username            string     `json:"username" gorm:"unique"`
	Email               string     `json:"email" gorm:"unique"`
	Password            string     `json:"password"`
	SpotifyAccessToken  string     `json:"-"`
	SpotifyRefreshToken string     `json:"-"`
	Playlists           []Playlist `json:"-"`
}

type Playlist struct {
	gorm.Model
	UserID uint
	Name   string
	Tracks []Track
}

type Track struct {
	gorm.Model
	SpotifyID  string
	PlaylistID uint
	Name       string
	Link       string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Models return one of every model the database must create..
func Models() []interface{} {
	return []interface{}{
		&Playlist{},
		&Track{},
	}
}
