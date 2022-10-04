package main

import (
	"net/http"
	"time"

	"github.com/guri-security/go-pipeline-test-api/routes"
)

var password = "senha_teste_sast_123"

func main() {

	server := &http.Server{
		Addr:              ":8888",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           routes.Load(),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
