package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/config"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/lyouthzzz/go-web-layout/internal/repo"
	"github.com/lyouthzzz/go-web-layout/internal/service"
	"github.com/lyouthzzz/go-web-layout/internal/usecase"
	"github.com/lyouthzzz/go-web-layout/pkg/app"
	"github.com/lyouthzzz/go-web-layout/pkg/gormx"
	"log"
)

func NewHttpServer(conf *config.Config) app.Lifecycle {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())

	apiRender := &render.APIRender{}

	mysqlDns := gormx.MysqlDns(conf.DataBase.User, conf.DataBase.Password, conf.DataBase.Host, conf.DataBase.Port, conf.DataBase.DbName, "")
	db, err := gormx.Connect(mysqlDns, true)
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userService := service.NewUserService(userUsecase)
	v1.RegisterUserServiceHTTPServer(userService, engine, apiRender)

	httpSvr := v1.NewHttpServer(conf.HttpServer.Host, conf.HttpServer.Port, engine)
	return httpSvr
}
