package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	models "github.com/epousa/ginFrameworkPractise/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAPIData(c *gin.Context) {
	var result models.ApodResponse

	//Get API data
	resp, err := http.Get(os.Getenv("APOD") + os.Getenv("NASA_KEY"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.IndentedJSON(resp.StatusCode, gin.H{"error": "Failed to fetch data"})
		return
	}

	//read API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	//Parse API JSON response to struct format
	if err := json.Unmarshal(body, &result); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response body"})
		log.Println(err) // Log the error for debugging purposes
		return
	}

	//Display API Data
	c.IndentedJSON(http.StatusOK, result)
}

// getAlbums responds with the list of all albums as JSON.
// handler func
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Users)
}

func SignUp(c *gin.Context) {
	var newUser models.User

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	// Add the new album to the slice.
	models.Users = append(models.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func SignIn(c *gin.Context) {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	for _, u := range models.Users {
		if u.Name == user.Name && u.Password == user.Password {
			c.IndentedJSON(http.StatusOK, gin.H{
				"msg": "Signed in successfully.",
			})
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}

func GetNoRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{})
}
