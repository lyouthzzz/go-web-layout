package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, &cmd{Code: -1, Data: nil, Message: err.Error()})
}

func (ar *APIRender) OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &cmd{Code: 0, Data: data, Message: "success"})
}
