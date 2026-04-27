package hillsystem

import (
	"cryptoapi/pkg/dictionary"
	"cryptoapi/pkg/matrix"
	"errors"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Encode(key, str string) (string, error) {
	keyRunes := []rune(key)
	strRunes := []rune(str)

	keyMatrix, vLen, err := matrix.ConvertToMatrix(keyRunes)
	if err != nil {
		return "", err
	}

	det, err := matrix.DeterminantMod(keyMatrix, dictionary.LenPrintableRunes)
	if err != nil {
		return "", err
	}
	if g, _, _ := matrix.ExtendedGCD(det, dictionary.LenPrintableRunes); g != 1 {
		return "", errors.New("ключ необратим: определитель не взаимно прост с модулем")
	}

	vectors, err := matrix.ConvertToVectors(strRunes, vLen)
	if err != nil {
		return "", err
	}

	var encodedRunes []rune
	for _, vec := range vectors {
		encVec, err := matrix.MultiplyMatrixVectorMod(keyMatrix, vec, dictionary.LenPrintableRunes)
		if err != nil {
			return "", err
		}
		for _, idx := range encVec {
			encodedRunes = append(encodedRunes, dictionary.KeyboardPrintableRunes[idx])
		}
	}

	return string(encodedRunes), nil
}

func (s *Service) Decode(key, str string) (string, error) {
	keyRunes := []rune(key)
	strRunes := []rune(str)

	keyMatrix, vLen, err := matrix.ConvertToMatrix(keyRunes)
	if err != nil {
		return "", err
	}

	invMatrix, err := matrix.InverseMatrixMod(keyMatrix, dictionary.LenPrintableRunes)
	if err != nil {
		return "", err
	}

	vectors, err := matrix.ConvertToVectors(strRunes, vLen)
	if err != nil {
		return "", err
	}

	var decodedRunes []rune
	for _, vec := range vectors {
		decVec, err := matrix.MultiplyMatrixVectorMod(invMatrix, vec, dictionary.LenPrintableRunes)
		if err != nil {
			return "", err
		}
		for _, idx := range decVec {
			decodedRunes = append(decodedRunes, dictionary.KeyboardPrintableRunes[idx])
		}
	}

	return string(decodedRunes), nil
}
