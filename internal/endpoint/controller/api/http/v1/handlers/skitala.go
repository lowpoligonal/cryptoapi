package handlers

import (
	skitalaService "cryptoapi/internal/domain/services/chiphers/skitala"
	"cryptoapi/pkg/response"
)

type SkitalaHandler struct {
	service *skitalaService.Service
}

func NewSkitalaHandler() *SkitalaHandler {
	return &SkitalaHandler{service: skitalaService.NewService()}
}

func (h *SkitalaHandler) Handle(mode string, req response.Request) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyInt, req.Text)
	}
	return h.service.Decode(req.KeyInt, req.Text)
}
