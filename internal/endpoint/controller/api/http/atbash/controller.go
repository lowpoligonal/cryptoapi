package atbash

import (
	"cryptoapi/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type atbashService interface {
	Transform(str string) (string, error)
}

type Controller struct {
	service atbashService
}

func NewController(s atbashService) *Controller {
	return &Controller{service: s}

}

func (h *Controller) Init(r *gin.RouterGroup) {
	atbash := r.Group("/atbash")
	{
		atbash.POST("/encode", h.encode)
		atbash.POST("/decode", h.decode)
	}

}

func (h *Controller) encode(c *gin.Context) {
	var req response.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Transform(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": res})
}

func (h *Controller) decode(c *gin.Context) {
	var req response.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Transform(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": res})
}
