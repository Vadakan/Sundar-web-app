package main

import (
	"fmt"
	"net/http"

	"github.com/Webapp-New/Handlers"
)

var (
	portNumber = ":8080"
)

// main is the main executable function
func main() {

	http.HandleFunc("/home", Handlers.Home)
	http.HandleFunc("/about", Handlers.About)

	fmt.Println(fmt.Sprintf("starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
