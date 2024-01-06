package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"os"
)

var flagConfig string

func init() {
	flag.StringVar(&flagConfig, "config", "configs/config.yaml", "config path, eg: --config config.yaml")
}

func newApp(logger log.Logger, grpcServer *grpc.Server, httpServer *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name("go-web-layout"),
		kratos.Version("0.0.1"),
		kratos.Logger(logger),
		kratos.Server(grpcServer, httpServer),
	)
}

func main() {
	flag.Parse()

	logger := log.With(log.NewStdLogger(os.Stdout))

	c := config.New(
		config.WithSource(
			file.NewSource(flagConfig),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
