package routes

import (
	"fmt"
	"net/http"
)

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func getApi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "getApi\n")
}

func Load() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/api", getApi)
}
