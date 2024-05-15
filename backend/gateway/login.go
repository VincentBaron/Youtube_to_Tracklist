package gateway

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"gorm.io/gorm"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type tokenInfo struct {
}

type User struct {
	gorm.Model
	UserID              string
	Username            string
	SpotifyAccessToken  string
	SpotifyRefreshToken string
	SpotifyExpiry       time.Time
}

var (
	auth = spotify.NewAuthenticator("http://localhost:8080/callback", spotify.ScopeUserReadPrivate)
	// Replace with your client ID and secret
	clientID     = "your-client-id"
	clientSecret = "your-client-secret"
	state        = "your-state"
)

func spotifyLoginHandler(c *gin.Context) {
	SpotifyClientID, ok := c.Get("SpotifyClientID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SpotifyClientID does not exist "})
		return

	}
	SpotifyClientID = SpotifyClientID.(string)
	SpotifyClientSecret, ok := c.Get("SpotifyClientSecret")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SpotifyClientSecret does not exist "})
		return
	}
	SpotifyClientSecret = SpotifyClientSecret.(string)
	auth.SetAuthInfo(SpotifyClientID.(string), SpotifyClientSecret.(string))
	url := auth.AuthURL("your-state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func callbackHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	token, err := auth.Token("your-state", c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// use the token to get an authenticated client
	client := auth.NewClient(token)

	// get the current user's ID
	user, err := client.CurrentUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get current user"})
		return
	}

	// save the token to the database
	db.Create(&User{
		UserID:              user.ID,
		Username:            user.DisplayName,
		SpotifyAccessToken:  token.AccessToken,
		SpotifyRefreshToken: token.RefreshToken,
		SpotifyExpiry:       token.Expiry,
	})

	c.JSON(http.StatusOK, gin.H{"status": "Success"})
}

// func signupHandler(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var form LoginForm
// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := User{UserID: form.Username}
// 	db.First(&user)

// 	if user.ID != 0 {
// 		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
// 		return
// 	}

// 	db.Create(&User{UserID: form.Username})

// 	c.JSON(http.StatusOK, gin.H{"status": "Success"})
// }

// func loginHandler(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var form LoginForm
// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var user User
// 	if err := db.Where("user_id = ?", form.Username).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": "Success"})
// }

// func storeTokenHandler(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var form tokenInfo
// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var user User
// 	if err := db.Where("user_id = ?", form.Username).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	user.SpotifyAccessToken = form.Password
// 	db.Save(&user)

// 	c.JSON(http.StatusOK, gin.H{"status": "Success"})
// }
