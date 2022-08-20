package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"user/pkg/log"
)

// New opens connection and ping database
func New(cfg Config, l log.Logger) (*sqlx.DB, error) {
	host := findHost(cfg.ConnectionString)
	l.Infof("open: connection to %q", host)

	db, err := open(cfg)
	if err == nil {
		return db, nil
	}

	for i := 0; i < cfg.MaxRetries; i++ {
		db, err = open(cfg)
		if err != nil {
			l.Errorf("retry: connection to %q: %v", host, err)
			time.Sleep(cfg.RetryDelay)
			continue
		}

		return db, nil
	}

	return nil, fmt.Errorf("failed: open connection to %q: %w\n", host, err)
}

func open(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, cfg.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
	if cfg.ConnMaxIdleNum > 0 {
		db.SetMaxIdleConns(cfg.ConnMaxIdleNum)
	}
	if cfg.ConnMaxOpenNum > 0 {
		db.SetMaxOpenConns(cfg.ConnMaxIdleNum)
	}

	return db, nil
}

func findHost(conn string) string {
	substring := "host="
	for _, value := range strings.Split(conn, " ") {
		if strings.Contains(value, substring) {
			return strings.TrimLeft(value, substring)
		}
	}

	return ""
}
