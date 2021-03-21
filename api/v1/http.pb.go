package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Middlewares struct {
	Session gin.HandlerFunc
}

type Servers struct {
	UserServer *UserServer
}

type HttpOption struct {
	Host        string
	Port        int
	Servers     *Servers
	Middlewares *Middlewares
	Logger      *zerolog.Logger
}

type HttpServer struct {
	host        string
	port        int
	engine      *gin.Engine
	servers     *Servers
	middlewares *Middlewares
	logger      *zerolog.Logger
}

func NewHttpServer(opt *HttpOption) *HttpServer {
	return &HttpServer{host: opt.Host, port: opt.Port, servers: opt.Servers, middlewares: opt.Middlewares, logger: opt.Logger}
}

func (svr *HttpServer) BuildRouter() error {
	engine := gin.New()

	apiV1 := engine.Group("/api/v1")
	apiV1.Use(gin.Recovery(), gin.Logger())

	sessionRequired := svr.middlewares.Session
	userSvr := svr.servers.UserServer

	apiV1.POST("/user/login", userSvr.Login)
	apiV1.POST("/user/logout", userSvr.Logout)
	apiV1.POST("/user/register", userSvr.CreateUser)

	apiV1.GET("/user/:id", sessionRequired, userSvr.GetUser)
	apiV1.PUT("/user/:id", sessionRequired, userSvr.UpdateUser)
	apiV1.DELETE("/user/:id", sessionRequired, userSvr.DeleteUser)

	svr.engine = engine
	return nil
}

func (svr *HttpServer) Start(ctx context.Context) error {
	return svr.engine.Run(fmt.Sprintf(fmt.Sprintf("%s:%d", svr.host, svr.port)))
}

func (svr *HttpServer) Stop(ctx context.Context) error {
	return nil
}
