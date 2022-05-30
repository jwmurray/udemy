package main

import (
	"fmt"
	"log"
	"net/http"
)

const LISTENING_PORT string = ":8000"

func log_on_error(err error) {
	if err != nil {
		log.Println("Error condition %s", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "<H1>Hello, there.</H1>")
	log_on_error(err)
	fmt.Println(fmt.Sprintf("The number of bytes written: %d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<H1>About page.</H1>")
	log_on_error(err)
}

func WardList(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<H1>Wardlist page.</H1>")
	log_on_error(err)
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/wardlist", WardList)

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(LISTENING_PORT, nil)

}
