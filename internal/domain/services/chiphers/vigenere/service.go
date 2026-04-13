package vigenere

import "cryptoapi/internal/domain/dictionary"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) Encode(key, str string) (string, error) {
	strRune := []rune(str)
	keyRune := []rune(key)
	var encoded []rune

	keyPos := 0

	for _, rn := range strRune {
		alph, idx := dictionary.FindSymbolInfo(rn)

		keyChar := keyRune[keyPos%len(keyRune)]
		_, shift := dictionary.FindSymbolInfo(keyChar)

		newIdx := (idx + shift) % len(dictionary.Dictionary[alph])
		encoded = append(encoded, dictionary.Dictionary[alph][newIdx])

		keyPos++
	}
	return string(encoded), nil
}

func (srv *Service) Decode(key, str string) (string, error) {
	strRune := []rune(str)
	keyRune := []rune(key)
	var decoded []rune

	keyPos := 0

	for _, rn := range strRune {
		alph, idx := dictionary.FindSymbolInfo(rn)
		alphLen := len(dictionary.Dictionary[alph])

		keyChar := keyRune[keyPos%len(keyRune)]
		_, shift := dictionary.FindSymbolInfo(keyChar)

		newIdx := (idx - shift + alphLen*2) % alphLen
		decoded = append(decoded, dictionary.Dictionary[alph][newIdx])

		keyPos++
	}
	return string(decoded), nil
}
