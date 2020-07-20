package persistence

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestSaveCandidateMutant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mutant := sqlmock.NewRows([]string{"dna", "is_mutant"})
	mock.ExpectQuery("SELECT (.+) FROM Candidate WHERE dna = \\?").WithArgs("").WillReturnRows(mutant)
	mock.ExpectExec("INSERT INTO Candidate").WithArgs("", true).WillReturnResult(sqlmock.NewResult(1, 1))
	dbApi := NewMysqlImp()
	candidate := GetNewCandidate("", true)
	err = dbApi.SaveCandidate(db, candidate)
	if err != nil {
		t.Errorf("An error '%s' was not expected when making data base operation", err)
	}
}

func TestSaveCandidateNotMutant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mutant := sqlmock.NewRows([]string{"dna", "is_mutant"}).AddRow("", 1)
	mock.ExpectQuery("SELECT (.+) FROM Candidate WHERE dna = \\?").WithArgs("").WillReturnRows(mutant)
	mock.ExpectExec("INSERT INTO Candidate").WithArgs("", true).WillReturnResult(sqlmock.NewResult(1, 1))
	dbApi := NewMysqlImp()
	candidate := GetNewCandidate("", true)
	err = dbApi.SaveCandidate(db, candidate)
	if err != nil {
		t.Errorf("An error '%s' was not expected when making data base operation", err)
	}
}

func TestStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rowsNotMutants := sqlmock.NewRows([]string{"count"}).AddRow(1)
	rowsMutants := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count").WillReturnRows(rowsNotMutants)
	mock.ExpectQuery("SELECT count").WillReturnRows(rowsMutants)
	dbApi := NewMysqlImp()
	stats, error := dbApi.GetStats(db)
	if error != nil {
		t.Errorf("An error '%s' was not expected when making data base operation", err)
	}
	if stats.GetCountHumans() != 1 || stats.GetCountHumans() != 1 {
		t.Errorf("Error, got %d, expected %d", stats.GetCountHumans(), 1)
	}
}
