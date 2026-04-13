package v1

import (
	"cryptoapi/internal/endpoint/controller/api/http/v1/handlers"
	"cryptoapi/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	chiphers map[string]handlers.ChipherHandler
}

func NewController() *Controller {
	return &Controller{chiphers: NewRegistry()}
}

func (h *Controller) Init(r *gin.RouterGroup) {
	chipherGroup := r.Group("/chipher/:chipher")
	{
		chipherGroup.POST("/encode", h.encode)
		chipherGroup.POST("/decode", h.decode)
	}
}

func (h *Controller) encode(c *gin.Context) {
	h.handle(c, "encode")
}

func (h *Controller) decode(c *gin.Context) {
	h.handle(c, "decode")
}

func (h *Controller) handle(c *gin.Context, mode string) {
	chipherName := c.Param("chipher")

	handler, ok := h.chiphers[chipherName]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported chipher"})
		return
	}

	var req response.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.Handle(mode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
