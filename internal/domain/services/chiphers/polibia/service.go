package polibia

import (
	"cryptoapi/internal/domain/dictionary"
	"errors"
	"unicode"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) Encode(str string) (string, error) {
	strRune := []rune(str)
	var encoded []rune

	var regMask []bool
	var rows []int
	var cols []int

	langFound, alphMap, alphMatrix := langFind(strRune)

	if !langFound {
		return str, errors.New("letters not found")
	}

	for _, ch := range strRune {
		chUp := unicode.ToUpper(ch)
		if coords, exists := alphMap[chUp]; exists {
			regMask = append(regMask, unicode.IsUpper(ch))
			rows = append(rows, coords[0])
			cols = append(cols, coords[1])
		}
	}
	coordStr := append(cols, rows...)

	var encodedChars []rune
	for i := 0; i < len(coordStr); i += 2 {
		col := coordStr[i]
		row := coordStr[i+1]
		outChar := alphMatrix[row][col]

		if regMask[i/2] {
			encodedChars = append(encodedChars, unicode.ToUpper(outChar))
		} else {
			encodedChars = append(encodedChars, unicode.ToLower(outChar))
		}
	}

	letterIndex := 0
	for _, char := range strRune {
		charUp := unicode.ToUpper(char)
		if _, exists := alphMap[charUp]; exists {
			if letterIndex < len(encodedChars) {
				encoded = append(encoded, encodedChars[letterIndex])
				letterIndex++
			}
		} else {
			encoded = append(encoded, char)
		}
	}

	return string(encoded), nil
}

func (srv *Service) Decode(str string) (string, error) {
	strRune := []rune(str)
	var decoded []rune

	var regMask []bool
	var coordStr []int

	langFound, alphMap, alphMatrix := langFind(strRune)

	if !langFound {
		return str, errors.New("letters not found")
	}

	for _, ch := range strRune {
		chUp := unicode.ToUpper(ch)
		if coords, exists := alphMap[chUp]; exists {
			regMask = append(regMask, unicode.IsUpper(ch))
			coordStr = append(coordStr, coords[1], coords[0])
		}
	}

	var decodedChars []rune
	cols := coordStr[:(len(coordStr) / 2)]
	rows := coordStr[(len(coordStr) / 2):]

	for i := 0; i < len(coordStr)/2; i++ {
		col := cols[i]
		row := rows[i]
		outChar := alphMatrix[row][col]

		if regMask[i] {
			decodedChars = append(decodedChars, unicode.ToUpper(outChar))
		} else {
			decodedChars = append(decodedChars, unicode.ToLower(outChar))
		}
	}

	letterIndex := 0
	for _, char := range strRune {
		charUp := unicode.ToUpper(char)
		if _, exists := alphMap[charUp]; exists {
			if letterIndex < len(decodedChars) {
				decoded = append(decoded, decodedChars[letterIndex])
				letterIndex++
			}
		} else {
			decoded = append(decoded, char)
		}
	}

	return string(decoded), nil
}

func langFind(strRune []rune) (bool, map[rune][2]int, [][]rune) {
	langFound := false
	var alphMatrix [][]rune
	var alphMap map[rune][2]int
	for _, char := range strRune {
		alph, _ := dictionary.FindSymbolInfo(char)
		switch alph {
		case "rusLow", "rusUp":
			alphMap = dictionary.RusMap
			alphMatrix = dictionary.RusMatrix
			langFound = true
		case "enLow", "enUp":
			alphMap = dictionary.EngMap
			alphMatrix = dictionary.EngMatrix
			langFound = true
		}
		if langFound {
			break
		}
	}
	return langFound, alphMap, alphMatrix
}
