package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// managing gin web server from tutorial https://go.dev/doc/tutorial/web-service-gin
// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s
// contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized
// field names – a style not as common in JSON.
// digimon represents data about a record digimon.
type digimon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Level          int    `json:"level"`
	Stage          string `json:"stage"`
	Attribute      string `json:"attribute"`
	PriorEvolution int    `json:"priorEvolution"`
	NextEvolution  int    `json:"nextEvolution"`
}

// digimons slice to seed record digimon data.
var digimons = []digimon{
	{ID: 1047, Name: "War Greymon (X-Antibody)", Level: 4, Stage: "Ultimate", Attribute: "Vaccine", PriorEvolution: 76, NextEvolution: 1309},
	{ID: 771, Name: "Lucemon", Level: 7, Stage: "Child", Attribute: "Vaccine", PriorEvolution: 833, NextEvolution: 879},
	{ID: 746, Name: "Punimon", Level: 5, Stage: "Baby I", Attribute: "Data", PriorEvolution: -1, NextEvolution: 1154},
}

// Example JSON
// {
//     "id": 1334,
//     "name": "Alphamon",
//     "level": 4,
//     "stage": "Ultimate",
//     "attribute": "Vaccine",
//     "priorEvolution": 1374,
//     "nextEvolution": 560
// }

// When the client makes a request at GET /digimons, you want to return all the digimons as JSON.
// getDigimons is a handler function that responds with the list of all digimons as JSON.
// you could have given this function any name – neither Gin nor Go require a particular function name format.
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more.
func getDigimons(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// The function’s first argument is the HTTP status code you want to send to the client.
	// The StatusOK constant from the net/http package to indicate 200 OK.
	//NOTE - You can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
	//The indented form is much easier to work with when debugging and the size difference is usually small.
	c.IndentedJSON(http.StatusOK, digimons)

}

// postDigimons adds a new digimon from JSON received in the request body.
func postDigimons(c *gin.Context) {
    var newDigimon digimon

    // Call BindJSON to bind the received JSON to
    // newDigimon.
    if err := c.BindJSON(&newDigimon); err != nil {
        return
    }

    // Add the new digimon to the slice.
    digimons = append(digimons, newDigimon)
	// Add a 201 status code to the response, along with JSON representing the digimon you added.
    c.IndentedJSON(http.StatusCreated, newDigimon)
}

// getDigimonByID locates the digimon whose ID value matches the id
// parameter sent by the client, then returns that digimon as a response.
func getDigimonByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of digimons, looking for a digimon whose ID value matches the parameter.
	// If it’s found, you serialize that digimon struct to JSON and return it as a response with a 200 OK HTTP code
    for _, d := range digimons {
        if d.ID == id {
            c.IndentedJSON(http.StatusOK, d)
            return
        }
    }

	// Return an HTTP 404 error with http.StatusNotFound if the digimon isn’t found.
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "digimon not found"})
}

func main() {
	// This sets up an association in which getDigimons handles requests to the /digimons endpoint path.
	// Initialize a Gin router using Default.
	router := gin.Default()
	// With Gin, you can associate a handler with an HTTP method-and-path combination. 
	// This way, you can separately route requests sent to a single path based on the method the client is using.
	// Use the GET function to associate the GET HTTP method and /digimons path with a handler function.
	router.GET("/digimons", getDigimons)
	// Associate the /digimons/:id path with the getDigimonByID function. In Gin, the colon preceding an item in the 
	// path signifies that the item is a path parameter.
	router.GET("/digimons/:id", getDigimonByID)
	// Associate the POST method at the /digimons path with the postDigimons function.
	router.POST("/digimons", postDigimons)
	
	// Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:3000")
}
