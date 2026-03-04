package main

import (
	"cryptoapi/internal/domain/services/atbash"
	"cryptoapi/internal/domain/services/skitala"
	atbashController "cryptoapi/internal/endpoint/controller/api/http/atbash"
	skitalaController "cryptoapi/internal/endpoint/controller/api/http/skitala"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	atbashSrv := atbash.NewService()
	atbashCtrl := atbashController.NewController(atbashSrv)

	skitalaSrv := skitala.NewService()
	skitalaCtrl := skitalaController.NewController(skitalaSrv)

	api := r.Group("/api/v1")
	atbashCtrl.Init(api)
	skitalaCtrl.Init(api)

	r.Run(":8080")
}
