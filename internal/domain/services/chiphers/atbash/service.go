package atbash

import "cryptoapi/internal/domain/dictionary"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) Transform(str string) (string, error) {
	var encrypt string
	for _, char := range str {
		alph, idx := dictionary.FindSymbolInfo(char)

		newIdx := len(dictionary.Dictionary[alph]) - 1 - idx
		encrypt += string(dictionary.Dictionary[alph][newIdx])
	}
	return encrypt, nil
}
