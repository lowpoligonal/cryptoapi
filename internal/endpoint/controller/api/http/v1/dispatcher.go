package v1

import "github.com/gin-gonic/gin"

type Controller interface {
	Init(*gin.RouterGroup)
}

type Dispatcher struct {
	ctrls []Controller
}

func NewDispatcher(ctrls ...Controller) *Dispatcher {
	return &Dispatcher{ctrls: ctrls}
}

func (d *Dispatcher) Init(r *gin.RouterGroup) {
	for _, ctrl := range d.ctrls {
		ctrl.Init(r)
	}
}
