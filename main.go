package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the Home page and 2 + 2 is %d", sum))
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Can't divide by Zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("This is the Divide page and %f / %f is %f", 100.0, 0.0, f))
}

func addValue(x, y int) int {
	return x + y
}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can't divide by Zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Start application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
