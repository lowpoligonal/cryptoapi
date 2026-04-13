package handlers

import (
	vigenereService "cryptoapi/internal/domain/services/chiphers/vigenere"
	"cryptoapi/pkg/response"
)

type VigenereHandler struct {
	service *vigenereService.Service
}

func NewVigenereHandler() *VigenereHandler {
	return &VigenereHandler{service: vigenereService.NewService()}
}

func (h *VigenereHandler) Handle(mode string, req response.Request) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyStr, req.Text)
	}
	return h.service.Decode(req.KeyStr, req.Text)
}
