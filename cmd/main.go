package main

import (
	"cryptoapi/internal/domain/services/chiphers/atbash"
	"cryptoapi/internal/domain/services/chiphers/caesar"
	"cryptoapi/internal/domain/services/chiphers/skitala"
	atbashController "cryptoapi/internal/endpoint/controller/api/http/atbash"
	caesarController "cryptoapi/internal/endpoint/controller/api/http/caesar"
	skitalaController "cryptoapi/internal/endpoint/controller/api/http/skitala"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	atbashSrv := atbash.NewService()
	atbashCtrl := atbashController.NewController(atbashSrv)

	skitalaSrv := skitala.NewService()
	skitalaCtrl := skitalaController.NewController(skitalaSrv)

	caesarSrv := caesar.NewService()
	caesarCtrl := caesarController.NewController(caesarSrv)

	api := r.Group("/api")
	atbashCtrl.Init(api)
	skitalaCtrl.Init(api)
	caesarCtrl.Init(api)

	r.Run(":8080")
}
