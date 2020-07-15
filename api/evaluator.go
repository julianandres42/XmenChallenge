package api

import (
	"regexp"
)

type MutantEvaluator interface {
	IsMutant() bool
}

type DnaSequenceEvaluator struct {
	Dna []string
}

func (evaluator DnaSequenceEvaluator) IsMutant() bool {
	return IsMutant(evaluator.Dna)
}

func IsMutant(dna []string) bool {
	if !validateAdn(dna) {
		return false
	}
	dnaMatrix := createMatrix(dna)
	return checkSequences(dnaMatrix)
}

func checkSequences(dna [][]string) bool {
	size := len(dna)
	totalSequences := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var rowSequences int
			rowSequences = findSequences(dna, i, j)
			totalSequences += rowSequences
			if totalSequences > 1 {
				return true
			}
		}
		if totalSequences > 1 {
			return true
		}
	}
	return totalSequences > 1
}

func findSequences(dna [][]string, i int, j int) int {
	totalSequences := 0
	if i+4 <= len(dna)-1 {
		totalSequences += findVerticalSequence(dna, i, j)
		if totalSequences > 1 {
			return totalSequences
		}
	}
	if j+4 <= len(dna)-1 {
		totalSequences += findHorizontalSequence(dna, i, j)
		if totalSequences > 1 {
			return totalSequences
		}
	}
	if i+4 <= len(dna)-1 && j+4 <= len(dna)-1 {
		totalSequences += findDiagonalSequence(dna, i, j)
		if totalSequences > 1 {
			return totalSequences
		}
	}
	return totalSequences
}

func findDiagonalSequence(dna [][]string, i int, j int) int {
	sequence := ""
	for index1, index2 := i, j; index1 < len(dna) && index2 < len(dna); index1, index2 = index1+1, index2+1 {
		sequence += dna[index1][index2]
	}
	return evaluateMatchSequence(sequence)
}

func findHorizontalSequence(dna [][]string, i int, j int) int {
	sequence := ""
	for index := j; index < len(dna[i]); index++ {
		sequence += dna[i][index]
	}
	return evaluateMatchSequence(sequence)
}

func findVerticalSequence(dna [][]string, i int, j int) int {
	sequence := ""
	for index := i; index < len(dna); index++ {
		sequence += dna[index][j]
	}
	return evaluateMatchSequence(sequence)
}

func evaluateMatchSequence(sequence string) int {
	reg := regexp.MustCompile("(?:AAAA|TTTT|CCCC|GGGG)(?:\\s+(?:AAAA|TTTT|CCCC|GGGG))*")
	sequencesFound := reg.FindAllString(sequence, -1)
	return len(sequencesFound)
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
	return isNxN(dna) && matchPattern(dna) && len(dna) >= 4
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
