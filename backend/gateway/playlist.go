package gateway

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zmb3/spotify"

	"github.com/gin-gonic/gin"
)

type TracklistResponse struct {
	Tracks []struct {
		TrackTitle string      `json:"trackTitle"`
		Timestamp  interface{} `json:"timestamp"`
	} `json:"tracks"`
}

type SpotifyAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type CreatePlaylistReq struct {
	URL         string `json:"url" binding:"required"`
	AccessToken string `json:"accessToken" binding:"required"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type SpotifyPlaylistResponse struct {
	// other fields...
	ID           string       `json:"id"`
	ExternalUrls ExternalUrls `json:"external_urls"`
}

type SpotifySearchResponse struct {
	Tracks struct {
		Items []struct {
			URI string `json:"uri"`
		} `json:"items"`
	} `json:"tracks"`
}

func CreatePlaylist(c *gin.Context) {

	log.Println("Creating playlist...")
	tmpClient, exists := c.Get("client")
	spotifyClient := tmpClient.(spotify.Client)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve Spotify client"})
		return
	}

	var body CreatePlaylistReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("GET", "https://www.1001tracklists.com/tracklist/1yvs6s7k/fred-again..-antro-juan-cdmx-mexico-2024-04-26.html", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", "guid=21041a5fa0ed8; shortscc=4")
	req.Header.Add("dnt", "1")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("sec-ch-ua", `"Chromium";v="123", "Not:A-Brand";v="8"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"macOS"`)
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tracks []string
	doc.Find(".tlpItem").Each(func(i int, s *goquery.Selection) {
		track := s.Find(".trackValue").Text()
		track = strings.TrimSpace(track)
		tracks = append(tracks, track)
	})

	// ...

	user, err := spotifyClient.CurrentUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Create a new playlist
	playlist, err := spotifyClient.CreatePlaylistForUser(user.ID, "New Playlist", "New playlist description", false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Create a new Playlist record
	// playlistDB := db.Create(&models.Playlist{
	// 	SpotifyID: string(playlist.ID),
	// 	UserID:    user.ID,
	// 	Name:      playlist.Name,
	// })

	// Search for the track URIs and add the tracks to the playlist
	for _, track := range tracks {
		results, err := spotifyClient.Search(track, spotify.SearchTypeTrack)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(results.Tracks.Tracks) > 0 {
			trackID := results.Tracks.Tracks[0].ID
			_, err = spotifyClient.AddTracksToPlaylist(playlist.ID, trackID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// // Create a new Track record
			// db.Create(&models.Track{
			// 	SpotifyID:  string(trackID),
			// 	PlaylistID: playlist.ID,
			// 	Name:       results.Tracks.Tracks[0].Name,
			// 	Link:       results.Tracks.Tracks[0].ExternalURLs["spotify"],
			// })
		}
	}

	c.JSON(http.StatusOK, gin.H{"playlistURI": playlist.URI})
}
