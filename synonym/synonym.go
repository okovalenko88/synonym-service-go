package main

import (
	"sync"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	Start()
}

func Start() {
	router := gin.Default()
	router.POST("/synonyms", postSynonyms)
	router.GET("/synonyms", getSynonyms)
	router.GET("/synonyms/:key", getSynonymsByKey)

	router.Run("localhost:8080")
}

// Data structures
type synonymsPost struct {
	Words []string `json:"words"`
}

var synonymsMap = make(map[string][]string)

type SafeMap struct {
	mu sync.Mutex
	sm map[string][]string
}

var safeSynonymsMap = SafeMap{sm: synonymsMap}

// Handlers
func getSynonyms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, synonymsMap)
}

func getSynonymsByKey(c *gin.Context) {
	key := c.Param("key")

	if synonyms, exists := synonymsMap[key]; exists {
		c.IndentedJSON(http.StatusOK, synonyms)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "synonym not found"})
}

func postSynonyms(c *gin.Context) {
	var newSynonyms synonymsPost

	if err := c.BindJSON(&newSynonyms); err != nil {
		return
	}

	var words = newSynonyms.Words

	updateSynonyms(&safeSynonymsMap, words)
}

// Functionality
func updateSynonyms(sm *SafeMap, words []string) {
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
		sm.safeMapInsert(word, synonymWords)
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

/*
 * Thread safe map insert
 */
func (sm *SafeMap) safeMapInsert(key string, values []string) {
	sm.mu.Lock()
	sm.sm[key] = values
	sm.mu.Unlock()
}
