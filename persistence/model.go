package persistence

import (
	"database/sql"
)

type DbApi interface {
	SaveCandidate(candidate *Candidate) error
	GetStats() (*Stats, error)
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

type MysqlImp struct {
	dataSource *sql.DB
}

func (imp *MysqlImp) SaveCandidate(candidate *Candidate) error {
	rows, err := imp.dataSource.Query("SELECT * FROM Candidate WHERE dna = ?", candidate.getDna())
	if err != nil {
		return err
	}
	if !rows.Next() {
		_, err = imp.dataSource.Exec("INSERT INTO Candidate (dna, is_mutant) VALUES (?, ?)", candidate.getDna(), candidate.getIsMutant())
		return err
	}
	return nil
}

func (imp *MysqlImp) GetStats() (*Stats, error) {
	var nMutants int
	err := imp.dataSource.QueryRow("SELECT count(*) FROM Candidate WHERE is_mutant = ?", true).Scan(&nMutants)
	if err != nil {
		return nil, err
	}
	var nHumans int
	err = imp.dataSource.QueryRow("SELECT count(*) FROM Candidate WHERE is_mutant = ?", false).Scan(&nHumans)
	if err != nil {
		return nil, err
	}
	return GetNewStats(nMutants, nHumans), nil
}

func NewMysqlImp(db *sql.DB) DbApi {
	return &MysqlImp{dataSource: db}
}
