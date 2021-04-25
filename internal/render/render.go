package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

type APIRender struct {
}

func (ar *APIRender) Error(c *gin.Context, err error) {
	err = errors.Cause(err)
	c.JSON(http.StatusOK, &cmd{Code: -1, Data: nil, Message: err.Error()})
}

func (ar *APIRender) OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &cmd{Code: 0, Data: data, Message: "success"})
}
