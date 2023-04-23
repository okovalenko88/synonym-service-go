package synonym

import (
	"sync"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/synonyms", getSynonyms)
	router.POST("/synonyms", postSynonyms)

	router.Run("localhost:8080")
}

type synonyms struct {
	Words []string `json:"words"`
}

// Store the synonyms into a map
var synonymsMap = make(map[string][]string)

// SafeMap is safe to use concurrently.
type SafeMap struct {
	mu sync.Mutex
	sm map[string][]string
}

var safeSynonymsMap = SafeMap{sm: synonymsMap}

func getSynonyms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, synonymsMap)
}

func postSynonyms(c *gin.Context) {
	var newSynonyms synonyms

	if err := c.BindJSON(&newSynonyms); err != nil {
		return
	}

	var words = newSynonyms.Words

	updateSynonyms(words)
}

func updateSynonyms(words []string) {
	key := words[0]
	_, exists := synonymsMap[words[0]]
	if exists {
		words = removeDuplicates(
			append(words, synonymsMap[key]...))
	}
	for i, word := range words {
		var wordsCopy = make([]string, len(words))
		copy(wordsCopy, words)
		synonymWords := append(wordsCopy[0:i], wordsCopy[i+1:]...)
		safeSynonymsMap.threadSafeMapInsert(word, synonymWords)
	}
}

func removeDuplicates(lst []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range lst {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func (sm *SafeMap) threadSafeMapInsert(key string, values []string) {
	sm.mu.Lock()
	// sm.sm[key] = append(sm.sm[key], values...)
	sm.sm[key] = values
	sm.mu.Unlock()
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
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
