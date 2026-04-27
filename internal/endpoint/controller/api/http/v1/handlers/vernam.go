package handlers

import (
	vernamService "cryptoapi/internal/domain/services/chiphers/vernam"
	"cryptoapi/pkg/response"
)

type VernamHandler struct {
	service *vernamService.Service
}

func NewVernamHandler() *VernamHandler {
	return &VernamHandler{service: vernamService.NewService()}
}

func (h *VernamHandler) Handle(mode string, req response.Response) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyStr, req.Text)
	}
	return h.service.Decode(req.KeyStr, req.Text)
}
