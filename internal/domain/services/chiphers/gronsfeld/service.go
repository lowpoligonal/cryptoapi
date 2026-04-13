package gronsfeld

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
		shift := int(keyRune[keyPos%len(keyRune)] - '0')
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
		shift := int(keyRune[keyPos%len(keyRune)] - '0')
		newIdx := (idx - shift + alphLen) % alphLen
		decoded = append(decoded, dictionary.Dictionary[alph][newIdx])

		keyPos++
	}
	return string(decoded), nil
}
