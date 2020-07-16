package app

import (
	"XmenChallenge/api"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	/*connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}*/
	a.Router = mux.NewRouter()
	a.Evaluator = api.NewDnaSequenceEvaluator()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/mutant", a.isMutant).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) isMutant(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var mutant Mutant
	json.Unmarshal(reqBody, &mutant)
	if mutant.Dna == nil || len(mutant.Dna) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	a.Evaluator.SetDna(mutant.Dna)
	if !a.Evaluator.IsMutant() {
		writer.WriteHeader(http.StatusForbidden)
		return
	}
}
