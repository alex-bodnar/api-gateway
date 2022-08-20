package database

import (
	"time"
)

type (
	// Config - configuration for postgreSQL database
	Config struct {
		ConnectionString   string        `yaml:"connection-string"`
		ConnMaxIdleNum     int           `yaml:"conn-max-idle-num"`
		ConnMaxOpenNum     int           `yaml:"conn-max-open-num"`
		Dialect            string        `yaml:"dialect"`
		Driver             string        `yaml:"driver"`
		MaxRetries         int           `yaml:"max-retries"`
		ConnMaxLifetime    time.Duration `yaml:"conn-max-lifetime"`
		RetryDelay         time.Duration `yaml:"retry-delay"`
		QueryTimeout       time.Duration `yaml:"query-timeout"`
		AutoMigrate        bool          `yaml:"auto-migrate"`
		MigrationDirectory string        `yaml:"migration-directory"`
	}
)
