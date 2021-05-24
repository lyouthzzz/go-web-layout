package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/go-web-layout/pkg/ecode"
	"github.com/pkg/errors"
)

type Render interface {
	Error(c *gin.Context, err error)
	OK(c *gin.Context, data interface{})
}

type cmd struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type EcodeRender struct {
}

func (ar *EcodeRender) Error(c *gin.Context, err error) {
	var se = &ecode.StatusError{}
	if errors.As(err, se) {
		c.JSON(se.HTTPStatus(), &cmd{Code: int(se.Code), Message: se.Reason})
		return
	}
	c.JSON(500, &cmd{Code: 500, Message: err.Error()})
}

func (ar *EcodeRender) OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &cmd{Code: 0, Data: data, Message: "success"})
}
