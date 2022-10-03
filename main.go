package main

import (
	"net/http"
	"time"

	"github.com/guri-security/go-pipline-test-api/routes"
)

func main() {
	routes.Load()

	server := &http.Server{
		Addr:              ":8888",
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
