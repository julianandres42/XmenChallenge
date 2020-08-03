package api

import (
	"regexp"
	"strings"
	"sync"
)

type MutantEvaluator interface {
	IsMutant() bool
	SetDna([]string)
}

type DnaSequenceEvaluator struct {
	Dna []string
}

func (evaluator *DnaSequenceEvaluator) IsMutant() bool {
	return IsMutant(evaluator.Dna)
}

func (evaluator *DnaSequenceEvaluator) SetDna(dna []string) {
	evaluator.Dna = dna
}

func NewDnaSequenceEvaluator() MutantEvaluator {
	return &DnaSequenceEvaluator{}
}

func IsMutant(dna []string) bool {
	if !validateAdn(dna) {
		return false
	}
	dnaMatrix := createMatrix(dna)
	sequencesFound := findDiagonalUpperSequence(dnaMatrix, 0)
	if sequencesFound > 1 {
		return true
	}
	var w sync.WaitGroup
	var mg sync.Mutex
	size := len(dnaMatrix)
	half := size / 2
	w.Add(6)
	go searchSequences(dnaMatrix, 0, half-1, findHorizontalSequence, &sequencesFound, &w, &mg)
	go searchSequences(dnaMatrix, half, size-1, findHorizontalSequence, &sequencesFound, &w, &mg)
	go searchSequences(dnaMatrix, 0, half-1, findVerticalSequence, &sequencesFound, &w, &mg)
	go searchSequences(dnaMatrix, half, size-1, findVerticalSequence, &sequencesFound, &w, &mg)
	go searchSequences(dnaMatrix, 1, size-4, findDiagonalUpperSequence, &sequencesFound, &w, &mg)
	go searchSequences(dnaMatrix, 1, size-4, findDiagonalLowerSequence, &sequencesFound, &w, &mg)
	w.Wait()
	return sequencesFound > 1
}

func searchSequences(dna [][]string,
	start, end int,
	sequenceCalculation func([][]string, int) int,
	sequencesFound *int,
	wg *sync.WaitGroup,
	m *sync.Mutex) {
	for i := start; i <= end; i++ {
		if *sequencesFound > 1 {
			wg.Done()
			return
		}
		sequences := 0
		sequences = sequenceCalculation(dna, i)
		m.Lock()
		*sequencesFound += sequences
		m.Unlock()
	}
	wg.Done()
}

func findDiagonalUpperSequence(dna [][]string, y int) int {
	sequence := ""
	for index1, index2 := 0, y; index1 < len(dna) && index2 < len(dna); index1, index2 = index1+1, index2+1 {
		sequence += dna[index1][index2]
	}
	return evaluateMatchSequence(sequence)
}

func findDiagonalLowerSequence(dna [][]string, x int) int {
	sequence := ""
	for index1, index2 := x, 0; index1 < len(dna) && index2 < len(dna); index1, index2 = index1+1, index2+1 {
		sequence += dna[index1][index2]
	}
	return evaluateMatchSequence(sequence)
}

func findHorizontalSequence(dna [][]string, i int) int {
	return evaluateMatchSequence(strings.Join(dna[i], ""))
}

func findVerticalSequence(dna [][]string, i int) int {
	sequence := ""
	for index := 0; index < len(dna); index++ {
		sequence += dna[index][i]
	}
	return evaluateMatchSequence(sequence)
}

func evaluateMatchSequence(sequence string) int {
	reg := regexp.MustCompile("(?:AAAA|TTTT|CCCC|GGGG)(?:\\s+(?:AAAA|TTTT|CCCC|GGGG))*")
	return len(reg.FindAllString(sequence, -1))
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
