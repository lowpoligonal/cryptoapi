package vernam

import (
	"encoding/hex"
	"errors"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Service struct{}

func NewService() *Service { return &Service{} }

func (s *Service) Encode(key, str string) (string, error) {
	strBytes, err := charmap.Windows1251.NewEncoder().Bytes([]byte(str))
	if err != nil {
		return "", errors.New("Текст содержит недопустимые символы")
	}
	keyBytes, err := charmap.Windows1251.NewEncoder().Bytes([]byte(key))
	if err != nil {
		return "", errors.New("Ключ содержит недопустимые символы")
	}

	keyPos := 0
	transformed := make([]byte, len(strBytes))
	for i, char := range strBytes {
		keyChar := keyBytes[keyPos%len(keyBytes)]
		transformed[i] = char ^ keyChar

		keyPos++
	}
	endString := hex.EncodeToString(transformed)

	return endString, nil
}

func (s *Service) Decode(key, str string) (string, error) {
	cleanHex := strings.ReplaceAll(strings.TrimSpace(str), " ", "")
	strBytes, err := hex.DecodeString(cleanHex)
	if err != nil {
		return "", errors.New("Ошибка кодирования")
	}

	keyBytes, err := charmap.Windows1251.NewEncoder().Bytes([]byte(key))
	if err != nil {
		return "", errors.New("Ключ содержит символы, которых нет в cp1251")
	}

	keyPos := 0
	transformed := make([]byte, len(strBytes))
	for i, char := range strBytes {
		keyChar := keyBytes[keyPos%len(keyBytes)]
		transformed[i] = char ^ keyChar

		keyPos++
	}

	decoded, err := charmap.Windows1251.NewDecoder().Bytes(transformed)
	if err != nil {
		return "", errors.New("Не удалось декодировать результат")
	}

	return string(decoded), nil
}
