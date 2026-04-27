package handlers

import (
	"cryptoapi/pkg/response"

	playfairService "cryptoapi/internal/domain/services/chiphers/playfair"
)

type PlayfairHandler struct {
	service *playfairService.Service
}

func NewPlayfairHandler() *PlayfairHandler {
	return &PlayfairHandler{service: playfairService.NewService()}
}

func (h *PlayfairHandler) Handle(mode string, req response.Response) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyStr, req.Text)
	}
	return h.service.Decode(req.KeyStr, req.Text)
}
