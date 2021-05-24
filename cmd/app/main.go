package main

import (
	"github.com/douyu/jupiter"
	gormx "github.com/douyu/jupiter/pkg/store/gorm"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/lyouthzzz/go-web-layout/internal/data"
	"github.com/lyouthzzz/go-web-layout/internal/server"
	"github.com/lyouthzzz/go-web-layout/internal/service"
	"github.com/lyouthzzz/go-web-layout/internal/usecase"
)

func newApp() (*jupiter.Application, func(), error) {
	app := &jupiter.Application{}
	// 初始化application
	if err := app.Startup(); err != nil {
		return nil, nil, err
	}

	logger := xlog.JupiterLogger

	// 初始化db
	db := gormx.StdConfig("example").Build()

	userData, err := data.NewData(db)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(userData, logger)
	userUsecase := usecase.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase, logger)

	httpServer, err := server.NewHTTPServer(userService)
	if err != nil {
		return nil, nil, err
	}
	grpcServer, err := server.NewGRPCServer(userService)
	if err != nil {
		return nil, nil, err
	}

	app.Serve(httpServer)
	app.Serve(grpcServer)

	cleanup := func() {
		_ = db.Close()
	}
	return app, cleanup, nil
}

func main() {
	app, cleanup, err := newApp()
	if err != nil {
		xlog.JupiterLogger.Fatal("init application failed.", xlog.FieldErr(err))
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		xlog.JupiterLogger.Fatal("run application failed.", xlog.FieldErr(err))
	}
}
