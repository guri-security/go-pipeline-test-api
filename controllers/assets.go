package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Asset struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
