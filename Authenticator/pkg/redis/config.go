package redis

import "time"

type (
	// Config - configuration for redis
	Config struct {
		ConnectionAddr string        `yaml:"connection-address"`
		DB             int           `yaml:"db"`
		Pass           string        `yaml:"pass"`
		MaxRetries     int           `yaml:"max-retries"`
		RetryDelay     time.Duration `yaml:"retry-delay"`
		QueryTimeout   time.Duration `yaml:"query-timeout"`
	}
)
