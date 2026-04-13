package v1

import "cryptoapi/internal/endpoint/controller/api/http/v1/handlers"

func NewRegistry() map[string]handlers.ChipherHandler {
	return map[string]handlers.ChipherHandler{
		"atbash":    handlers.NewAtbashHandler(),
		"caesar":    handlers.NewCaesarHandler(),
		"gronsfeld": handlers.NewGronsfeldHandler(),
		"polibia":   handlers.NewPolibiaHandler(),
		"skitala":   handlers.NewSkitalaHandler(),
		"vigenere":  handlers.NewVigenereHandler(),
	}
}
