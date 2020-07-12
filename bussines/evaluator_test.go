package bussines

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

func TestFindHorizontalSequence(t *testing.T) {
	adn := [][]string{{"A", "A", "A", "A", "C"}, {"A", "T", "C", "G", "T"}, {"A", "T", "C", "E", "E"}, {"A", "T", "C", "F", "F"}}
	indexFound := make(map[string]bool)
	var found bool
	indexFoundExpected := map[string]bool{"00": true, "01": true, "02": true, "03": true}
	found, indexFound = findHorizontalSequence(adn, indexFound, 0, 0)
	if !found {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	sizeIndexesFound := len(indexFound)
	sizeIndexesExpected := len(indexFoundExpected)
	if len(indexFoundExpected) != len(indexFound) {
		t.Errorf("Expected %d, got %d", sizeIndexesExpected, sizeIndexesFound)
	}
	for k, v := range indexFound {
		if v != indexFoundExpected[k] {
			t.Errorf("Expected true for the key %s", k)
		}
	}
}

func TestFindVerticalSequence(t *testing.T) {
	adn := [][]string{{"A", "A", "A", "A", "C"}, {"A", "T", "C", "G", "T"}, {"A", "T", "C", "E", "E"}, {"A", "T", "C", "F", "F"}}
	indexFound := make(map[string]bool)
	var found bool
	indexFoundExpected := map[string]bool{"00": true, "10": true, "20": true, "30": true}
	found, indexFound = findVerticalSequence(adn, indexFound, 0, 0)
	if !found {
		t.Errorf("Sequence not found in indexes 0 , 0")
	}
	sizeIndexesFound := len(indexFound)
	sizeIndexesExpected := len(indexFoundExpected)
	if len(indexFoundExpected) != len(indexFound) {
		t.Errorf("Expected %d, got %d", sizeIndexesExpected, sizeIndexesFound)
	}
	for k, v := range indexFound {
		if v != indexFoundExpected[k] {
			t.Errorf("Expected true for the key %s", k)
		}
	}
}
