package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerSignIn)
	http.HandleFunc("/callback", handlerCallback)

	fmt.Println("Listen here: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
