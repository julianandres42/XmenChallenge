package bussines

import (
	"regexp"
	"strconv"
)

func IsMutant(dna []string) bool {
	if !validateAdn(dna) {
		return false
	}
	dnaMatrix := createMatrix(dna)
	return checkSequences(dnaMatrix)
}

func checkSequences(dna [][]string) bool {
	var indexFound map[string]bool
	size := len(dna)
	totalSequences := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if !indexFound[string(i)+string(j)] {
				var rowSequences int
				rowSequences, indexFound = findSequences(dna, indexFound, i, j)
				totalSequences += rowSequences
				if totalSequences == 2 {
					return true
				}
			}
		}
		if totalSequences == 2 {
			return true
		}
	}
	return totalSequences >= 2
}

func findSequences(dna [][]string, indexFound map[string]bool, i int, j int) (int, map[string]bool) {
	totalSequences := 0
	foundSequence, indexUsed := findVerticalSequence(dna, indexFound, i, j)
	if foundSequence {
		totalSequences++
	}
	foundSequence, indexUsed = findHorizontalSequence(dna, indexUsed, i, j)
	if foundSequence {
		totalSequences++
	}
	if totalSequences == 2 {
		return totalSequences, indexUsed
	}
	foundSequence, indexUsed = findDiagonalUpSequence(dna, indexUsed, i, j)
	if foundSequence {
		totalSequences++
	}
	if totalSequences == 2 {
		return totalSequences, indexUsed
	}
	foundSequence, indexUsed = findDiagonalDownSequence(dna, indexUsed, i, j)
	if foundSequence {
		totalSequences++
	}
	if totalSequences == 2 {
		return totalSequences, indexUsed
	}
	return totalSequences, indexUsed
}

func findDiagonalDownSequence(dna [][]string, used map[string]bool, i int, j int) (bool, map[string]bool) {
	return true, used
}

func findDiagonalUpSequence(dna [][]string, used map[string]bool, i int, j int) (bool, map[string]bool) {
	return true, used
}

func findHorizontalSequence(dna [][]string, used map[string]bool, i int, j int) (bool, map[string]bool) {
	if dna[i][j] == dna[i][j+1] && dna[i][j] == dna[i][j+2] && dna[i][j] == dna[i][j+3] {
		used[strconv.Itoa(i)+strconv.Itoa(j)] = true
		used[strconv.Itoa(i)+strconv.Itoa(j+1)] = true
		used[strconv.Itoa(i)+strconv.Itoa(j+2)] = true
		used[strconv.Itoa(i)+strconv.Itoa(j+3)] = true
		return true, used
	}
	return false, used
}

func findVerticalSequence(dna [][]string, used map[string]bool, i int, j int) (bool, map[string]bool) {
	if dna[i][j] == dna[i+1][j] && dna[i][j] == dna[i+2][j] && dna[i][j] == dna[i+3][j] {
		used[strconv.Itoa(i)+strconv.Itoa(j)] = true
		used[strconv.Itoa(i+1)+strconv.Itoa(j)] = true
		used[strconv.Itoa(i+2)+strconv.Itoa(j)] = true
		used[strconv.Itoa(i+3)+strconv.Itoa(j)] = true
		return true, used
	}
	return false, used
}

func createMatrix(dna []string) [][]string {
	dnaMatrix := make([][]string, len(dna))
	for i, element := range dna {
		reg := regexp.MustCompile("")
		dnaMatrix[i] = reg.Split(element, -1)
	}
	return dnaMatrix
}

func validateAdn(dna []string) bool {
	return isNxN(dna) && matchPattern(dna)
}

func isNxN(dna []string) bool {
	size := len(dna)
	for _, element := range dna {
		if len(element) != size {
			return false
		}
	}
	return true
}

func matchPattern(dna []string) bool {
	reg := regexp.MustCompile("^[ATCG]+$")
	for _, element := range dna {
		if !reg.MatchString(element) {
			return false
		}
	}
	return true
}
