package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(response http.ResponseWriter, request *http.Request){
	fmt.Println("Hello " + request.URL.Path)
}
