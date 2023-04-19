package synonym

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/synonyms", getSynonyms)
	router.POST("/albums", postAlbums)
	router.POST("/synonyms", postSynonyms)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type synonyms struct {
	Words []string `json:"words"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Store the synonyms in a map
var synonymsMap = make(map[string][]string)

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getSynonyms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, synonymsMap)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func postSynonyms(c *gin.Context) {
	var newSynonyms synonyms

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newSynonyms); err != nil {
		return
	}

	var words = newSynonyms.Words

	for i, word := range words {
		var wordsCopy = make([]string, len(words))
		copy(wordsCopy, words)
		synonymsMap[word] = append(wordsCopy[0:i], wordsCopy[i+1:]...)
	}
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
