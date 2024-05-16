package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/VincentBaron/youtube_to_tracklist/backend/initializers"
	"github.com/VincentBaron/youtube_to_tracklist/backend/models"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		spotifyToken := &oauth2.Token{AccessToken: user.SpotifyAccessToken}
		client := spotify.Authenticator{}.NewClient(spotifyToken)
		c.Set("client", client)

		// Attach the request
		c.Set("user", user)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// func SpotifyAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the user from the context (you'll need to set this in another middleware)
// 		user, exists := c.Get("user")
// 		if !exists {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		// Check if the user has a Spotify access token
// 		if user.SpotifyAccessToken == "" {
// 			// If not, redirect them to the Spotify authentication flow
// 			url := "https://accounts.spotify.com/authorize" +
// 				"?response_type=code" +
// 				"&client_id=" + url.QueryEscape(clientID) +
// 				"&scope=" + url.QueryEscape("user-read-private user-read-email") +
// 				"&redirect_uri=" + url.QueryEscape("http://localhost:8080/callback")
// 			c.Redirect(http.StatusTemporaryRedirect, url)
// 			c.Abort()
// 			return
// 		}

// 		// Check if the access token is expired
// 		if isAccessTokenExpired(user.SpotifyAccessToken) {
// 			// If it is, use the refresh token to get a new access token
// 			newAccessToken, err := refreshAccessToken(user.SpotifyRefreshToken)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh access token"})
// 				c.Abort()
// 				return
// 			}

// 			// Update the user's access token in the database
// 			user.SpotifyAccessToken = newAccessToken
// 			initializers.DB.Save(&user)
// 		}

// 		// If the user has a valid Spotify access token, continue to the next handler
// 		c.Next()
// 	}
// }
