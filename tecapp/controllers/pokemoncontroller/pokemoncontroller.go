package pokemoncontroller

import (
	"net/http"
	"strconv"
	"tecapp/models/digimon"

	"github.com/gin-gonic/gin"
)

// When the client makes a request at GET /digimons, you want to return all the digimons as JSON.
// getDigimons is a handler function that responds with the list of all digimons as JSON.
// you could have given this function any name – neither Gin nor Go require a particular function name format.
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more.
func GetDigimons(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// The function’s first argument is the HTTP status code you want to send to the client.
	// The StatusOK constant from the net/http package to indicate 200 OK.
	//NOTE - You can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
	//The indented form is much easier to work with when debugging and the size difference is usually small.
	c.IndentedJSON(http.StatusOK, digimon.Digimons.Get())

}

// postDigimons adds a new digimon from JSON received in the request body.
func PostDigimons(c *gin.Context) {
	newDigimon := digimon.Digimons.New()

	// Call BindJSON to bind the received JSON to
	// newDigimon.
	if err := c.BindJSON(&newDigimon); err != nil {
		return
	}

	// Add a 201 status code to the response, along with JSON representing the digimon you added.
	c.IndentedJSON(http.StatusCreated, digimon.Digimons.Save(newDigimon))
}

// getDigimonByID locates the digimon whose ID value matches the id
// parameter sent by the client, then returns that digimon as a response.
func GetDigimonByID(c *gin.Context) {
	id, iderr := strconv.Atoi(c.Param("id"))

	// If there’s an error in parsing the `id` parameter as an integer, return a 400 Bad Request error.
	if iderr == nil {
		// Loop over the list of digimons, looking for a digimon whose ID value matches the parameter.
		// If it’s found, you serialize that digimon struct to JSON and return it as a response with a 200 OK HTTP code
		for _, d := range digimon.Digimons.Get() {
			if d.ID == id {
				c.IndentedJSON(http.StatusOK, d)
				return
			}
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	// Return an HTTP 404 error with http.StatusNotFound if the digimon isn’t found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "digimon not found"})
}
