package main

import (
	"log"
	"net/http"

	"github.com/Irl-Felix/simple-tcp-server/app"
)

func main() {
	http.HandleFunc("/", app.HealthCheckHandler)
	http.HandleFunc("/whoami", app.ClientinfoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
