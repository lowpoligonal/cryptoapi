package handlers

import (
	"cryptoapi/pkg/response"

	hillsystemService "cryptoapi/internal/domain/services/chiphers/hillsystem"
)

type HillSystemHandler struct {
	service *hillsystemService.Service
}

func NewHillSystemHandler() *HillSystemHandler {
	return &HillSystemHandler{service: hillsystemService.NewService()}
}

func (h *HillSystemHandler) Handle(mode string, req response.Response) (string, error) {
	if mode == "encode" {
		return h.service.Encode(req.KeyStr, req.Text)
	}
	return h.service.Decode(req.KeyStr, req.Text)
}
