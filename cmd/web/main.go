package main

import (
	"fmt"
	"net/http"

	"github.com/jwmurray/udemy/pkg/handlers"
)

const portNumber = ":8000"

// var Buf = new(byte.buffer)

func main() {
	http.HandleFunc("/about/", handlers.About)
	http.HandleFunc("/wardlist/", handlers.WardList)
	http.HandleFunc("/callings/", handlers.CallingList)
	http.HandleFunc("/", handlers.Home)

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(portNumber, nil)
}
