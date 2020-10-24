package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var serviceName = os.Getenv("SERVICE_NAME")
var pathPrefix = os.Getenv("PATH_PREFIX")

func main() {
	if serviceName == "" {
		fmt.Println("SERVICE_NAME is required")
		os.Exit(1)
	}
	if pathPrefix == "" {
		fmt.Println("PATH_PREFIX is required")
		os.Exit(1)
	}

	http.HandleFunc("/health", health)
	http.HandleFunc(fmt.Sprintf("/%s/fubar", pathPrefix), fubar)

	fmt.Println("server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy af on %s", serviceName)
}

func fubar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "things are fubar on %s", serviceName)
}
