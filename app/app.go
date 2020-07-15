package app

import (
	"XmenChallenge/api"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type App struct {
	Router    *mux.Router
	DB        *sql.DB
	Evaluator api.MutantEvaluator
}

type Mutant struct {
	Dna []string `json:"dna"`
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.Evaluator = api.DnaSequenceEvaluator{}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.isMutant).Methods("POST")
}

func (a *App) Run(addr string) {

}

func (a *App) isMutant(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var mutant Mutant
	json.Unmarshal(reqBody, &mutant)
	if mutant.Dna == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if !api.IsMutant(mutant.Dna) {
		writer.WriteHeader(http.StatusForbidden)
		return
	}
}
