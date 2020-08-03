package api

import (
	"sync"
	"testing"
)

func TestCreateRightMatrix(t *testing.T) {
	adn := []string{"ATCG", "ATCG", "ATCG", "ATCG"}
	size := len(adn)
	matrix := createMatrix(adn)
	if len(matrix) != size {
		t.Errorf("Expected %d, the size of the matrix must be the same as array", size)
	}
	for _, element := range matrix {
		if len(element) != size {
			t.Errorf("Expected %d, the size of each element must be the same as array", size)
		}
	}
}

func TestCreateWrongMatrix(t *testing.T) {
	adn := []string{"ATCGGG", "ATCGGGG", "ATC", "ATCGGGG", "ATGGGGG"}
	size := len(adn)
	matrix := createMatrix(adn)
	if len(matrix) != size {
		t.Errorf("Expected %d, the size of the matrix must be the same as array", size)
	}
	for _, element := range matrix {
		if len(element) == size {
			t.Errorf("Expected %d, the size of each element must be different as array", size)
		}
	}
}

func TestValidateRightAdn(t *testing.T) {
	adn := []string{"ATCG", "ATCG", "ATCG", "ATCG"}
	size := len(adn)
	if !validateAdn(adn) {
		t.Errorf("The size must be %d and the strings only can have the characters A, T, C, G", size)
	}
}

func TestValidateWrongSizeAdn(t *testing.T) {
	adn := []string{"ATCGG", "ATCG", "ATCGE", "ATCGF"}
	size := len(adn)
	if validateAdn(adn) {
		t.Errorf("The size must be %d and the strings only can have the characters A, T, C, G", size)
	}
}

func TestValidateWrongSequenceAdn(t *testing.T) {
	adn := []string{"ATCG", "ATCG", "ATCE", "ATCF"}
	size := len(adn)
	if validateAdn(adn) {
		t.Errorf("The size must be %d and the strings only can have the characters A, T, C, G", size)
	}
}

func TestEvaluateMatchSequence(t *testing.T) {
	sequence := "TTTTCAAAA"
	found := evaluateMatchSequence(sequence)
	if found != 2 {
		t.Errorf("Got %d, expected %d", found, 2)
	}
	sequence = "TTAAGGEE"
	found = evaluateMatchSequence(sequence)
	if found != 0 {
		t.Errorf("Got %d, expected %d", found, 2)
	}
}

func TestFindHorizontalSequence(t *testing.T) {
	dna := [][]string{{"A", "A", "A", "A", "C"}}
	found := findHorizontalSequence(dna, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	dna = [][]string{{"T", "T", "T", "T", "C", "A", "A", "A", "A"}}
	found = findHorizontalSequence(dna, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestFindVerticalSequence(t *testing.T) {
	dna := [][]string{{"A"}, {"A"}, {"A"}, {"A"}, {"C"}}
	found := findVerticalSequence(dna, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	dna = [][]string{{"T"}, {"T"}, {"T"}, {"T"}, {"C"}, {"A"}, {"A"}, {"A"}, {"A"}}
	found = findVerticalSequence(dna, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestFindDiagonalUpper(t *testing.T) {
	dna := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "T", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "T", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	found := findDiagonalUpperSequence(dna, 1)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestFindDiagonalLower(t *testing.T) {
	dna := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"T", "A", "T", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "T", "A", "T", "G"},
		{"C", "C", "C", "T", "T", "A"},
		{"T", "C", "A", "C", "A", "G"},
	}
	found := findDiagonalLowerSequence(dna, 1)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestSearchHorizontalSequences(t *testing.T) {
	var w sync.WaitGroup
	var mg sync.Mutex
	dna := [][]string{
		{"A", "T", "A", "C", "G", "A"},
		{"T", "T", "T", "T", "G", "C"},
		{"A", "A", "A", "A", "G", "T"},
		{"A", "T", "A", "A", "T", "G"},
		{"C", "C", "A", "T", "T", "A"},
		{"T", "C", "A", "C", "A", "G"},
	}
	sequencesFound := 0
	w.Add(1)
	searchSequences(dna, 1, 2, findHorizontalSequence, &sequencesFound, &w, &mg)
	w.Wait()
	if sequencesFound != 2 {
		t.Errorf("Sequence not found in indexes 1 , 2")
	}
}

func TestSearchVerticalSequences(t *testing.T) {
	var w sync.WaitGroup
	var mg sync.Mutex
	dna := [][]string{
		{"A", "T", "A", "C", "G", "A"},
		{"T", "T", "A", "T", "G", "C"},
		{"A", "T", "A", "A", "G", "T"},
		{"A", "T", "A", "A", "T", "G"},
		{"C", "C", "A", "T", "T", "A"},
		{"T", "C", "A", "C", "A", "G"},
	}
	sequencesFound := 0
	w.Add(1)
	searchSequences(dna, 1, 2, findVerticalSequence, &sequencesFound, &w, &mg)
	w.Wait()
	if sequencesFound != 2 {
		t.Errorf("Sequence not found in indexes 1 , 2")
	}
}

func TestSearchUpperDiagonalSequences(t *testing.T) {
	var w sync.WaitGroup
	var mg sync.Mutex
	dna := [][]string{
		{"A", "T", "A", "C", "G", "A"},
		{"T", "T", "T", "A", "G", "C"},
		{"A", "T", "A", "T", "A", "T"},
		{"A", "T", "A", "A", "T", "A"},
		{"C", "C", "A", "T", "T", "T"},
		{"T", "C", "A", "C", "A", "G"},
	}
	sequencesFound := 0
	w.Add(1)
	searchSequences(dna, 1, 2, findDiagonalUpperSequence, &sequencesFound, &w, &mg)
	w.Wait()
	if sequencesFound != 2 {
		t.Errorf("Sequence not found in indexes 1 , 2")
	}
}

func TestSearchLowerDiagonalSequences(t *testing.T) {
	var w sync.WaitGroup
	var mg sync.Mutex
	dna := [][]string{
		{"A", "T", "A", "C", "G", "A"},
		{"T", "T", "T", "A", "G", "C"},
		{"A", "T", "A", "A", "C", "T"},
		{"A", "A", "T", "A", "T", "A"},
		{"C", "C", "A", "T", "T", "T"},
		{"T", "C", "A", "A", "T", "G"},
	}
	sequencesFound := 0
	w.Add(1)
	searchSequences(dna, 1, 2, findDiagonalLowerSequence, &sequencesFound, &w, &mg)
	w.Wait()
	if sequencesFound != 2 {
		t.Errorf("Sequence not found in indexes 1 , 2")
	}
}

//func TestIsMutant(t *testing.T) {
//	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
//	evaluator := DnaSequenceEvaluator{}
//	evaluator.Dna = dna
//	isMutant := evaluator.IsMutant()
//	if !isMutant {
//		t.Errorf("Got not mutant")
//	}
//	dna = []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
//	evaluator.Dna = dna
//	isMutant = evaluator.IsMutant()
//	if isMutant {
//		t.Errorf("Got mutant")
//	}
//}
