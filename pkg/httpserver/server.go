package httpserver

import (
	"context"
	atbashService "cryptoapi/internal/domain/services/chiphers/atbash"
	caesarService "cryptoapi/internal/domain/services/chiphers/caesar"
	polibiaService "cryptoapi/internal/domain/services/chiphers/polibia"
	skitalaService "cryptoapi/internal/domain/services/chiphers/skitala"
	v1 "cryptoapi/internal/endpoint/controller/api/http/v1"
	"cryptoapi/internal/endpoint/controller/api/http/v1/atbash"
	"cryptoapi/internal/endpoint/controller/api/http/v1/caesar"
	"cryptoapi/internal/endpoint/controller/api/http/v1/polibia"
	"cryptoapi/internal/endpoint/controller/api/http/v1/skitala"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	router := gin.Default()
	router.Static("/static", "./web/static")

	router.GET("/", func(c *gin.Context) {
		c.File("./web/static/index.html")
	})

	atbashCtrl := atbash.NewController(atbashService.NewService())
	skitalaCtrl := skitala.NewController(skitalaService.NewService())
	polibiaCtrl := polibia.NewController(polibiaService.NewService())
	caesarCtrl := caesar.NewController(caesarService.NewService())

	dispatcher := v1.NewDispatcher(
		atbashCtrl,
		skitalaCtrl,
		polibiaCtrl,
		caesarCtrl)

	api := router.Group("/api")
	dispatcher.Init(api)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}

func (s *Server) Run(addr ...string) error {
	if len(addr) > 0 {
		s.httpServer.Addr = addr[0]
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
