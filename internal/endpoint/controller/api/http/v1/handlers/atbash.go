package handlers

import (
	atbashService "cryptoapi/internal/domain/services/chiphers/atbash"
	"cryptoapi/pkg/response"
)

type AtbashHandler struct {
	service *atbashService.Service
}

func NewAtbashHandler() *AtbashHandler {
	return &AtbashHandler{service: atbashService.NewService()}
}

func (h *AtbashHandler) Handle(mode string, req response.Request) (string, error) {
	return h.service.Transform(req.Text)
}
