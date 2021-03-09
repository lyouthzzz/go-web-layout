package main

import (
	"context"
	"github.com/jinzhu/configor"
	"github.com/lyouthzzz/go-web-layout/internal/config"
	"github.com/lyouthzzz/go-web-layout/internal/server"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var cfgFp string

	app := cli.NewApp()

	app.Name = "web-layout"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config,c",
			Value:       "configs/config.yaml",
			Destination: &cfgFp,
		},
	}
	app.Action = func(c *cli.Context) error {
		var (
			cfg config.Config
			err error
		)
		if _, err = os.Stat(cfgFp); err != nil {
			return err
		}

		if err = configor.Load(&cfg, cfgFp); err != nil {
			return err
		}

		httpSvr := server.NewHttpServer(&cfg)
		return httpSvr.Start(context.TODO())
	}

	panic(app.Run(os.Args))
}
