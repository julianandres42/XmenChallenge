package persistence

import (
	"database/sql"
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

func (s *Stats) GetCountMutants() int {
	return s.countMutants
}

func (s *Stats) GetCountHumans() int {
	return s.countHumans
}

func (s *Stats) GetRatio() float32 {
	return float32(s.countMutants) / float32(s.countHumans)
}

func GetNewStats(countMutants, countHumans int) *Stats {
	return &Stats{countHumans: countHumans, countMutants: countMutants}
}

type MysqlImp struct{}

func (imp *MysqlImp) SaveCandidate(db *sql.DB, candidate *Candidate) error {
	var nRecords int
	err := db.QueryRow("SELECT count(*) FROM Candidate WHERE dna = ?", candidate.getDna()).Scan(&nRecords)
	if err != nil {
		return err
	}
	if nRecords == 0 {
		_, err = db.Exec("INSERT INTO Candidate (dna, is_mutant) VALUES (?, ?)", candidate.getDna(), candidate.getIsMutant())
		return err
	}
	return nil
}

func (imp *MysqlImp) GetStats(db *sql.DB) (*Stats, error) {
	var nMutants int
	err := db.QueryRow("SELECT count(*) FROM Candidate WHERE is_mutant = ?", true).Scan(&nMutants)
	if err != nil {
		return nil, err
	}
	var nHumans int
	err = db.QueryRow("SELECT count(*) FROM Candidate WHERE is_mutant = ?", false).Scan(&nHumans)
	if err != nil {
		return nil, err
	}
	return GetNewStats(nMutants, nHumans), nil
}

func NewMysqlImp() DbApi {
	return &MysqlImp{}
}
