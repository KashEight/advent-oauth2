package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerSignIn)
	http.HandleFunc("/callback", handlerCallback)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
