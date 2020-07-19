package app

import (
	"XmenChallenge/persistence"
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type T struct {
}

func (T) IsMutant() bool {
	return true
}

func (T) SetDna(dna []string) {}

type T2 struct {
}

func (T2) IsMutant() bool {
	return false
}

func (T2) SetDna(dna []string) {}

type T3 struct {
}

func (T3) SaveCandidate(db *sql.DB, candidate *persistence.Candidate) error {
	return nil
}

func (T3) GetStats(db *sql.DB) (*persistence.Stats, error) {
	return persistence.GetNewStats(1, 1), nil
}

func TestInitialize(t *testing.T) {
	app := App{Evaluator: T{}, DbApi: T3{}}
	app.Initialize("", "", "", "")
}

func TestIsMutant(t *testing.T) {
	app := App{Evaluator: T{}, DbApi: T3{}}
	requestBody, err := json.Marshal(map[string][]string{"dna": {"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}})

	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	app.isMutant(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
}

func TestIsNotMutant(t *testing.T) {
	app := App{Evaluator: T2{}, DbApi: T3{}}
	requestBody, err := json.Marshal(map[string][]string{"dna": {"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}})

	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	app.isMutant(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
}

func TestBadRequest(t *testing.T) {
	app := App{Evaluator: T2{}}
	requestBody, err := json.Marshal(map[string][]string{"dna": {}})
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	app.isMutant(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
}

func TestStats(t *testing.T) {
	app := App{Evaluator: T2{}, DbApi: T3{}}
	stats := Stats{}
	requestBody, err := json.Marshal(map[string][]string{"dna": {}})
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	req, err := http.NewRequest("GET", "/stats", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	app.stats(rec, req)
	json.Unmarshal(rec.Body.Bytes(), &stats)
	if stats.CountHuman != 1 || stats.CountMutant != 1 || stats.Ratio != 1 {
		t.Errorf("expected %d, got: %d , %d,  got %d, %d, got %f", 1, stats.CountMutant, 1, stats.CountHuman, 1, stats.Ratio)
	}
}
