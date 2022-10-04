package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guri-security/go-pipeline-test-api/controllers"
)

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func getApi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "getApi\n")
}

func Load() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", getRoot)
	router.HandleFunc("/api", getApi)
	router.HandleFunc("/assets", controllers.CreateAsset).Methods("POST")

	return router
}
