package main

// managing gin web server from tutorial https://go.dev/doc/tutorial/web-service-gin
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}