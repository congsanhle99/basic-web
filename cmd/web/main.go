package main

import (
	"fmt"
	"net/http"

	"github.com/tsawler/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Start application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
