package app

import (
	"XmenChallenge/api"
	"XmenChallenge/persistence"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router    *mux.Router
	DB        *sql.DB
	Evaluator api.MutantEvaluator
	DbApi     persistence.DbApi
}

type Mutant struct {
	Dna []string `json:"dna"`
}

func (a *App) Initialize(user, password, host, dbname string) {
	a.EstablishDataBaseConnection(user, password, host, dbname)
	a.Router = mux.NewRouter()
	a.Evaluator = api.NewDnaSequenceEvaluator()
	a.DbApi = persistence.NewMysqlImp()
	a.initializeRoutes()
}

func (a *App) EstablishDataBaseConnection(user, password, host, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/mutant", a.isMutant).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
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
	isMutant := a.Evaluator.IsMutant()
	candidate := persistence.GetNewCandidate(strings.Join(mutant.Dna, ","), isMutant)
	err := a.DbApi.SaveCandidate(a.DB, candidate)
	if err != nil {
		print(fmt.Sprint("error in data base &s"), err)
	}
	if isMutant {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusForbidden)
	}
}
