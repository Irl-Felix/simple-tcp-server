package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy"))
	fmt.Println("Health check endpoint hit")
}
func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/", healthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
