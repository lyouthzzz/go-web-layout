package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/go-web-layout/pkg/app"
)

type HttpServer struct {
	host   string
	port   int
	engine *gin.Engine
}

func NewHttpServer(host string, port int, engine *gin.Engine) app.Lifecycle {
	return &HttpServer{host: host, port: port, engine: engine}
}

func (httpSvr *HttpServer) Start(ctx context.Context) error {
	return httpSvr.engine.Run(fmt.Sprintf(fmt.Sprintf("%s:%d", httpSvr.host, httpSvr.port)))
}

func (httpSvr *HttpServer) Stop(ctx context.Context) error {
	return nil
}
