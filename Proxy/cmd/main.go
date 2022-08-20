package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"proxy/internal/app"
	"proxy/internal/config"
)

var (
	appName       = "proxy-microservice"
	version       string
	commit        string
	tag           string
	date          string
	fortuneCookie string
)

func main() {
	cfgPath := flag.String("c", config.DefaultPath, "configuration file")
	flag.Parse()

	app.New(
		app.Meta{
			Info: app.Info{
				AppName:       appName,
				Tag:           tag,
				Version:       version,
				Commit:        commit,
				Date:          date,
				FortuneCookie: fortuneCookie,
			},
			ConfigPath: *cfgPath,
		},
	).Run(registerGracefulHandle())
}

func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
