// main
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// set the router
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("views/*")

	//define main digimon route
	router.GET("/digimon", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "TecApp"},
		)
	})

	//Start and run the server
	router.Run("localhost:3000")
}
