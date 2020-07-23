package api

import "testing"

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
	adn := [][]string{{"A", "A", "A", "A", "C"}}
	found := findHorizontalSequence(adn[0], 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	adn = [][]string{{"T", "T", "T", "T", "C", "A", "A", "A", "A"}}
	found = findHorizontalSequence(adn[0], 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestFindVerticalSequence(t *testing.T) {
	adn := [][]string{{"A"}, {"A"}, {"A"}, {"A"}, {"C"}}
	found := findVerticalSequence(adn, 0, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	adn = [][]string{{"T"}, {"T"}, {"T"}, {"T"}, {"C"}, {"A"}, {"A"}, {"A"}, {"A"}}
	found = findVerticalSequence(adn, 0, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestFindDiagonal(t *testing.T) {
	adn := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	found := findDiagonalSequence(adn, 0, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	adn = [][]string{
		{"A", "T", "G", "C", "G", "A", "A", "T", "C", "G"},
		{"C", "A", "G", "T", "G", "C", "A", "T", "C", "G"},
		{"T", "T", "A", "T", "G", "T", "A", "T", "C", "G"},
		{"A", "G", "A", "A", "G", "G", "A", "T", "C", "G"},
		{"C", "C", "C", "C", "T", "A", "A", "T", "C", "G"},
		{"T", "C", "A", "C", "T", "G", "A", "T", "C", "G"},
		{"T", "C", "A", "C", "T", "G", "T", "T", "C", "G"},
		{"T", "C", "A", "C", "T", "G", "T", "T", "C", "G"},
		{"T", "C", "A", "C", "T", "G", "T", "T", "T", "G"},
		{"T", "C", "A", "C", "T", "G", "T", "T", "T", "T"},
	}
	found = findDiagonalSequence(adn, 0, 0)
	if found < 1 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}

}

func TestFindSequences(t *testing.T) {
	adn := [][]string{
		{"A", "A", "A", "A", "G", "A"},
		{"A", "G", "G", "T", "G", "C"},
		{"A", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	sequencesFound := findSequences(adn, 0, 0)
	if sequencesFound != 2 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	adn = [][]string{
		{"A", "A", "G", "T", "G", "A"},
		{"A", "A", "G", "T", "G", "C"},
		{"A", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	sequencesFound = findSequences(adn, 0, 0)
	if sequencesFound != 2 {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
}

func TestCheckSequences(t *testing.T) {
	adn := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	sequencesFound := checkSequences(adn)
	if !sequencesFound {
		t.Errorf("Sequence not found in dna")
	}
}

func TestIsMutant(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	evaluator := DnaSequenceEvaluator{}
	evaluator.Dna = dna
	isMutant := evaluator.IsMutant()
	if !isMutant {
		t.Errorf("Got not mutant")
	}
	dna = []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
	evaluator.Dna = dna
	isMutant = evaluator.IsMutant()
	if isMutant {
		t.Errorf("Got mutant")
	}
}
