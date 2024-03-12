package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /json/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "json get request")
	})

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe("127.0.0.1:"+port, mux)
}
