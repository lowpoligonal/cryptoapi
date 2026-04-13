package polibia

import (
	"cryptoapi/internal/domain/dictionary"
	"errors"
	"unicode"
)

const (
	RusLang = iota
	EngLang
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) Encode(key int, str string) (string, error) {
	strRune := []rune(str)
	var encoded []rune

	var regMask []bool
	var rows []int
	var cols []int

	alphMap, alphMatrix, err := getLang(key)
	if err != nil {
		return "", err
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

func (srv *Service) Decode(key int, str string) (string, error) {
	strRune := []rune(str)
	var decoded []rune

	var regMask []bool
	var coordStr []int

	alphMap, alphMatrix, err := getLang(key)
	if err != nil {
		return "", err
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

func getLang(key int) (map[rune][2]int, [][]rune, error) {
	switch key {
	case RusLang:
		return dictionary.RusMap, dictionary.RusMatrix, nil
	case EngLang:
		return dictionary.EngMap, dictionary.EngMatrix, nil
	default:
		return nil, nil, errors.New("unsupported language key")
	}
}
