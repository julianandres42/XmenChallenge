package persistence

import (
	"database/sql"
	"errors"
)

type DbApi interface {
	SaveCandidate(db *sql.DB, candidate *Candidate) error
	GetStats(db *sql.DB) (*Stats, error)
}

type Candidate struct {
	dna      string
	isMutant bool
}

func GetNewCandidate(dna string, isMutant bool) *Candidate {
	return &Candidate{dna: dna, isMutant: isMutant}
}

func (c *Candidate) getDna() string {
	return c.dna
}

func (c *Candidate) getIsMutant() bool {
	return c.isMutant
}

type Stats struct {
	countMutants int
	countHumans  int
}

func (s *Stats) getRatio() float32 {
	return float32(s.countMutants) / float32(s.countHumans)
}

func GetNewStats(countMutants, countHumans int) *Stats {
	return &Stats{countHumans: countHumans, countMutants: countMutants}
}

type MysqlImp struct{}

func (imp *MysqlImp) SaveCandidate(db *sql.DB, candidate *Candidate) error {
	_, err := db.Exec("INSERT INTO Candidate (dna, is_mutant) VALUES (?, ?)", candidate.getDna(), candidate.getIsMutant())
	return err
}

func (imp *MysqlImp) GetStats(db *sql.DB) (*Stats, error) {
	return nil, errors.New("")
}

func NewMysqlImp() DbApi {
	return &MysqlImp{}
}
