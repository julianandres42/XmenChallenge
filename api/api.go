package api

import (
	"XmenChallenge/bussines"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Mutant struct {
	Dna []string `json:"dna"`
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/mutant", isMutant).Methods("POST")
	http.ListenAndServe(":8080", myRouter)
}

func isMutant(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var mutant Mutant
	json.Unmarshal(reqBody, &mutant)
	bussines.IsMutant(mutant.Dna)
}
