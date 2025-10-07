package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter , r *http.Request) {
    fmt.Fprintf(w , "Home Page")
}

func wsEndpoints(w http.ResponseWriter , r *http.Request) {
	  fmt.Fprintf(w , "Web Socket Endpoints")
}

func setUpRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoints)
}

func main() {
	setUpRoutes()
}
