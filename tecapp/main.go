// main
package main

import (
	"net/http"
	"tecapp/controllers/digimoncontroller"
	"tecapp/controllers/pokemoncontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	// set the router
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("views/*")

	//define main digimon route
	router.GET("/hello", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "TecApp", "IsWelcome": true},
		)
	})

	//Digimons Routes
	//define digimons HTML routes
	//route to get main digimons view
	router.GET("/digimons", digimoncontroller.Index)

	//define digimons JSON routes
	// Use the GET function to associate the GET HTTP method and /digimons path with a handler function.
	router.GET("/digimons/json", digimoncontroller.GetDigimonsJSON)
	// Associate the /digimons/:id path with the getDigimonByID function. In Gin, the colon preceding an item in the
	// path signifies that the item is a path parameter.
	router.GET("/digimons/json/:id", digimoncontroller.GetDigimonJSONByID)
	// Associate the POST method at the /digimons path with the postDigimons function.
	router.POST("/digimons/json", digimoncontroller.PostDigimon)

	//Pokemons Routes
	//define pokemons HTML routes
	//route to get main pokemons view
	router.GET("/pokemons", pokemoncontroller.Index)
	//route to get a specific pokemon view by id
	router.GET("/pokemons/:id", pokemoncontroller.Index)

	//define digimons JSON routes
	// Use the GET function to associate the GET HTTP method and /pokemons path with a handler function.
	router.GET("/pokemons/json", pokemoncontroller.GetDigimonsJSON)
	// Associate the /digimons/:id path with the getPokemonByID function. In Gin, the colon preceding an item in the
	// path signifies that the item is a path parameter.
	router.GET("/pokemons/json/:id", pokemoncontroller.GetDigimonJSONByID)
	// Associate the POST method at the /pokemons path with the postDigimons function.
	router.POST("/pokemons/json", pokemoncontroller.PostDigimon)

	//Start and run the server
	router.Run("localhost:3000")
}
