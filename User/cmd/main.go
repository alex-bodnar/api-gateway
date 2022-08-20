package main

import (
	"context"
	"embed"
	"flag"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/stdlib"

	"user/internal/app"
	"user/internal/config"
)

var (
	appName       = "user-microservice"
	version       string
	commit        string
	tag           string
	date          string
	fortuneCookie string
)

//go:embed dbschema/migrations
var dbMigrationFS embed.FS

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
	).WithMigrationFS(dbMigrationFS).Run(registerGracefulHandle())
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
