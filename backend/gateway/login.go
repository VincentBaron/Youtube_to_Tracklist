package gateway

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/VincentBaron/youtube_to_tracklist/backend/initializers"
	"github.com/VincentBaron/youtube_to_tracklist/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req Body
	var payload models.User

	if c.Bind(&payload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	// Create the user
	user := models.User{
		Email:    payload.Email,
		Username: payload.Username,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
	}

	state := strconv.FormatUint(uint64(user.ID), 10) // Convert the user's ID to a string
	// Redirect the user to the Spotify authorization page
	url := "https://accounts.spotify.com/authorize?response_type=code" +
		"&client_id=" + initializers.Conf.SpotifyClientID +
		"&scope=" + "playlist-modify-private" +
		"&redirect_uri=" + initializers.Conf.SpotifyRedirectURL +
		"&state=" + state

	log.Println(url)

	// // Respond
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func CallbackHandler(c *gin.Context) {
	// Get the code from the query parameters
	code := c.Query("code")

	// Prepare the data for the POST request
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", initializers.Conf.SpotifyRedirectURL)
	data.Set("client_id", initializers.Conf.SpotifyClientID)
	data.Set("client_secret", initializers.Conf.SpotifyClientSecret)

	// Create a new request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Set the content type to application/x-www-form-urlencoded
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Parse the response body
	var tokenResponse models.TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response body"})
		return
	}

	// Get the state parameter from the query parameters
	state := c.Query("state")

	// Convert the state parameter to a uint
	userID, err := strconv.ParseUint(state, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse state parameter"})
		return
	}

	// Update the user record with the access token and refresh token
	user := models.User{}
	initializers.DB.First(&user, userID)
	user.SpotifyAccessToken = tokenResponse.AccessToken
	user.SpotifyRefreshToken = tokenResponse.RefreshToken
	initializers.DB.Save(&user)

	// Log the response body
	log.Println(string(body))
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173")
}

func Login(c *gin.Context) {
	// Get email & pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up for requested user
	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in password with saved users password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	// user.(models.User).Email    -->   to access specific data

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
