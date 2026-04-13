package handlers

import (
	polibiaService "cryptoapi/internal/domain/services/chiphers/polibia"
	"cryptoapi/pkg/response"
)

type PolibiaHandler struct {
	service *polibiaService.Service
}

func NewPolibiaHandler() *PolibiaHandler {
	return &PolibiaHandler{service: polibiaService.NewService()}
}

func (h *PolibiaHandler) Handle(mode string, req response.Request) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyInt, req.Text)
	}
	return h.service.Decode(req.KeyInt, req.Text)
}
