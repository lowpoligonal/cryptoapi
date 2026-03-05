package caesar

import (
	"cryptoapi/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type caesarService interface {
	Decode(mKey int, str string) (string, error)
	Encode(mKey int, str string) (string, error)
}

type Controller struct {
	service caesarService
}

func NewController(s caesarService) *Controller {
	return &Controller{service: s}

}

func (h *Controller) Init(r *gin.RouterGroup) {
	caesar := r.Group("/caesar")
	{
		caesar.POST("/encode", h.encode)
		caesar.POST("/decode", h.decode)
	}
}

func (h *Controller) encode(c *gin.Context) {
	var req response.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Encode(req.Key, req.Text)
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

	res, err := h.service.Decode(req.Key, req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": res})
}
