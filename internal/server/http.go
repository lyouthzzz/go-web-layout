package server

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lyouthzzz/framework/pkg/app"
	"github.com/lyouthzzz/framework/pkg/gormx"
	v1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/api/v1/middleware/auth"
	"github.com/lyouthzzz/go-web-layout/internal/auth/authenticator"
	"github.com/lyouthzzz/go-web-layout/internal/auth/store"
	"github.com/lyouthzzz/go-web-layout/internal/config"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/lyouthzzz/go-web-layout/internal/repo"
	"github.com/lyouthzzz/go-web-layout/internal/service"
	"github.com/lyouthzzz/go-web-layout/internal/usecase"
	"github.com/rs/zerolog/log"
)

func NewHttpServer(conf *config.Config) app.Lifecycle {
	mysqlDns := gormx.MysqlDns(conf.DataBase.User, conf.DataBase.Password, conf.DataBase.Host, conf.DataBase.Port, conf.DataBase.DbName, "")
	db, err := gormx.Connect(mysqlDns, true)
	if err != nil {
		log.Fatal().Err(err)
	}

	rdb := redis.NewClient(&redis.Options{Addr: conf.Redis.Address, Password: conf.Redis.Password, DB: conf.Redis.DB})
	authN := authenticator.NewTokenAuthN(store.NewRDBStore(rdb, "session", time.Hour))

	sessionRequired := auth.Middleware(authN)

	apiRender := &render.APIRender{}

	userRepository := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	userService := service.NewUserService(userUsecase, authN)
	userServer := v1.NewUserHTTPServer(userService, apiRender)

	httpSvr := v1.NewHttpServer(&v1.HttpOption{
		Host:        conf.HttpServer.Host,
		Port:        conf.HttpServer.Port,
		Servers:     &v1.Servers{UserServer: userServer},
		Middlewares: &v1.Middlewares{Session: sessionRequired},
		Logger:      &log.Logger,
	})
	if err := httpSvr.BuildRouter(); err != nil {
		log.Fatal().Err(err)
	}
	return httpSvr
}
