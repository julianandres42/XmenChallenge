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
)

type App struct {
	DB        *sql.DB
	Evaluator api.MutantEvaluator
	DbApi     persistence.DbApi
}

type Mutant struct {
	Dna []string `json:"dna"`
}

type Stats struct {
	CountMutant int     `json:"count_mutant_dna"`
	CountHuman  int     `json:"count_human_dna"`
	Ratio       float32 `json:"ratio"`
}

func GetNewStats(countMutant, countHuman int, ratio float32) *Stats {
	return &Stats{CountMutant: countMutant, CountHuman: countHuman, Ratio: ratio}
}

func (a *App) Initialize(user, password, host, dbname, machine string) {
	a.EstablishDataBaseConnection(user, password, host, dbname, machine)
	a.Evaluator = api.NewDnaSequenceEvaluator()
	a.DbApi = persistence.NewMysqlImp(a.DB)
}

func (a *App) EstablishDataBaseConnection(user, password, host, dbname, machine string) {
	dbPatternConnection := ""
	if machine == "local" {
		dbPatternConnection = "%s:%s@tcp(%s)/%s"
	} else if machine == "cloud" {
		dbPatternConnection = "%s:%s@unix(/cloudsql/%s)/%s"
	}
	dbURI := fmt.Sprintf(dbPatternConnection, user, password, host, dbname)
	a.DB, _ = sql.Open("mysql", dbURI)
}

func (a *App) IsMutant(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Print("Getting is mutant")
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
	err := a.DbApi.SaveCandidate(candidate)
	if err != nil {
		fmt.Fprintf(writer, "error in data base connection")
		log.Print("Error in database connection")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("Success saving candidate")
	if isMutant {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusForbidden)
	}
}

func (a *App) Stats(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Print("Getting stats")
	dbStats, err := a.DbApi.GetStats()
	if err != nil {
		fmt.Fprintf(writer, "error in data base connection")
		log.Print("Error in database connection")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("Success getting stats")
	jsonStats := GetNewStats(dbStats.GetCountMutants(), dbStats.GetCountHumans(), dbStats.GetRatio())
	response, _ := json.Marshal(jsonStats)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
