package main

import (
	"fmt"

	"synonym.com/synonym"
)

func main() {
	// Get a greeting message and print it.
	message := synonym.Hello("Gladys")
	fmt.Println(message)
}

// func main() {
// 	router := gin.Default()
// 	router.GET("/albums", synonym.getAlbums)

// 	router.Run("localhost:8080")
// }
