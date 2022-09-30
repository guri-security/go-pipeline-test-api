package main

import (
	"net/http"

	"github.com/guri-security/go-pipline-test-api/routes"
)

func main() {
	routes.Load()
	http.ListenAndServe(":8888", nil)
}
