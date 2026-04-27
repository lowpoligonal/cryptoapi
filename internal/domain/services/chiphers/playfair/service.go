package playfair

import (
	"cryptoapi/pkg/dictionary"
	"math"
	"strings"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type cipherMatrix struct {
	matrix [][]rune
	size   int
	lookup map[rune][2]int
}

func initMatrix(key string) *cipherMatrix {
	baseRunes := append([]rune{dictionary.FillerRune}, dictionary.KeyboardPrintableRunes...)

	totalNeeded := len(baseRunes)
	size := int(math.Ceil(math.Sqrt(float64(totalNeeded))))

	used := make(map[rune]bool)
	var combined []rune

	for _, r := range key {
		if !used[r] {
			used[r] = true
			combined = append(combined, r)
		}
	}

	for _, r := range baseRunes {
		if !used[r] {
			used[r] = true
			combined = append(combined, r)
		}
	}

	matrix := make([][]rune, size)
	lookup := make(map[rune][2]int)
	for i := 0; i < size; i++ {
		matrix[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			idx := i*size + j
			var char rune
			if idx < len(combined) {
				char = combined[idx]
			} else {
				char = rune(idx + 1000)
			}
			matrix[i][j] = char
			lookup[char] = [2]int{i, j}
		}
	}

	return &cipherMatrix{matrix: matrix, size: size, lookup: lookup}
}

func (s *Service) Encode(key, text string) (string, error) {
	cm := initMatrix(key)
	runes := []rune(text)
	var prepared []rune

	for i := 0; i < len(runes); i++ {
		prepared = append(prepared, runes[i])
		if i+1 < len(runes) {
			if runes[i] == runes[i+1] {
				prepared = append(prepared, dictionary.FillerRune)
			} else {
				prepared = append(prepared, runes[i+1])
				i++
			}
		}
	}
	if len(prepared)%2 != 0 {
		prepared = append(prepared, dictionary.FillerRune)
	}
	return transform(prepared, cm, 1), nil
}

func (s *Service) Decode(key, text string) (string, error) {
	cm := initMatrix(key)
	decodedRunes := []rune(transform([]rune(text), cm, -1))

	var result []rune
	for i := 0; i < len(decodedRunes); i++ {
		if decodedRunes[i] == dictionary.FillerRune {

			if (i > 0 && i < len(decodedRunes)-1 && decodedRunes[i-1] == decodedRunes[i+1]) || i == len(decodedRunes)-1 {
				continue
			}
		}
		result = append(result, decodedRunes[i])
	}
	return string(result), nil
}

func transform(data []rune, cm *cipherMatrix, dir int) string {
	var res strings.Builder
	for i := 0; i < len(data); i += 2 {
		p1, ok1 := cm.lookup[data[i]]
		p2, ok2 := cm.lookup[data[i+1]]

		if !ok1 || !ok2 {
			res.WriteRune(data[i])
			res.WriteRune(data[i+1])
			continue
		}

		r1, c1, r2, c2, sz := p1[0], p1[1], p2[0], p2[1], cm.size
		if r1 == r2 {
			res.WriteRune(cm.matrix[r1][(c1+sz+dir)%sz])
			res.WriteRune(cm.matrix[r2][(c2+sz+dir)%sz])
		} else if c1 == c2 {
			res.WriteRune(cm.matrix[(r1+sz+dir)%sz][c1])
			res.WriteRune(cm.matrix[(r2+sz+dir)%sz][c2])
		} else {
			res.WriteRune(cm.matrix[r1][c2])
			res.WriteRune(cm.matrix[r2][c1])
		}
	}
	return res.String()
}
