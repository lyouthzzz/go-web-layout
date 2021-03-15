package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	v1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/config"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/lyouthzzz/go-web-layout/internal/repo"
	"github.com/lyouthzzz/go-web-layout/internal/service"
	"github.com/lyouthzzz/go-web-layout/internal/usecase"
	"github.com/lyouthzzz/go-web-layout/pkg/app"
	"github.com/lyouthzzz/go-web-layout/pkg/gormx"
	"log"
	"time"
)

func NewHttpServer(conf *config.Config) app.Lifecycle {
	engine := gin.New()

	apiV1 := engine.Group("/api/v1")
	apiV1.Use(gin.Recovery(), gin.Logger())

	apiRender := &render.APIRender{}

	mysqlDns := gormx.MysqlDns(conf.DataBase.User, conf.DataBase.Password, conf.DataBase.Host, conf.DataBase.Port, conf.DataBase.DbName, "")
	db, err := gormx.Connect(mysqlDns, true)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{Addr: conf.Redis.Address, Password: conf.Redis.Password, DB: conf.Redis.DB})

	sessionRepo := repo.NewSessionRepo(rdb)
	sessionUsecase := usecase.NewSessionUsecase(sessionRepo, time.Hour)

	userRepository := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	userService := service.NewUserService(userUsecase, sessionUsecase)
	v1.RegisterUserServiceHTTPServer(userService, apiV1, apiRender)

	httpSvr := v1.NewHttpServer(conf.HttpServer.Host, conf.HttpServer.Port, engine)
	return httpSvr
}
