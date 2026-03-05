package caesar

import "cryptoapi/internal/domain/dictionary"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) Encode(key int, str string) (string, error) {
	strRune := []rune(str)
	var encoded []rune

	for _, rn := range strRune {
		alph, idx := dictionary.FindSymbolInfo(rn)
		newIdx := (idx + key) % len(dictionary.Dictionary[alph])
		encoded = append(encoded, dictionary.Dictionary[alph][newIdx])
	}
	return string(encoded), nil
}

func (srv *Service) Decode(key int, str string) (string, error) {
	strRune := []rune(str)
	var decoded []rune

	for _, rn := range strRune {
		alph, idx := dictionary.FindSymbolInfo(rn)
		alphLen := len(dictionary.Dictionary[alph])
		newIdx := (idx - key) % alphLen
		if newIdx < 0 {
			newIdx += alphLen
		}
		decoded = append(decoded, dictionary.Dictionary[alph][newIdx])
	}
	return string(decoded), nil
}
