package handlers

import "cryptoapi/pkg/response"

type ChipherHandler interface {
	Handle(mode string, req response.Response) (string, error)
}
