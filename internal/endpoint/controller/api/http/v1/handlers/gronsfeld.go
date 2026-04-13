package handlers

import (
	gronsfeldService "cryptoapi/internal/domain/services/chiphers/gronsfeld"
	"cryptoapi/pkg/response"
)

type GronsfeldHandler struct {
	service *gronsfeldService.Service
}

func NewGronsfeldHandler() *GronsfeldHandler {
	return &GronsfeldHandler{service: gronsfeldService.NewService()}
}

func (h *GronsfeldHandler) Handle(mode string, req response.Request) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyStr, req.Text)
	}
	return h.service.Decode(req.KeyStr, req.Text)
}
