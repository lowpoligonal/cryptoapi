package handlers

import (
	caesarService "cryptoapi/internal/domain/services/chiphers/caesar"
	"cryptoapi/pkg/response"
)

type CaesarHandler struct {
	service *caesarService.Service
}

func NewCaesarHandler() *CaesarHandler {
	return &CaesarHandler{service: caesarService.NewService()}
}

func (h *CaesarHandler) Handle(mode string, req response.Request) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyInt, req.Text)
	}
	return h.service.Decode(req.KeyInt, req.Text)
}
