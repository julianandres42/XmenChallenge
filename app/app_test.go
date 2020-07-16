package app

import (
	"bytes"
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

func TestIsMutant(t *testing.T) {
	app := App{Evaluator: T{}}
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
	app := App{Evaluator: T2{}}
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
