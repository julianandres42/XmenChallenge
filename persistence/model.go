package persistence

import (
	"database/sql"
	"errors"
)

type Candidate struct {
	Dna      string
	IsMutant bool
}

func (c *Candidate) createCandidate(db *sql.DB) error {
	return errors.New("Not implemented")
}

type Stats struct {
	CountMutants int
	CountHumans  int
}

func (s *Stats) getStats(db *sql.DB) error {
	return errors.New("Not implemented")
}
