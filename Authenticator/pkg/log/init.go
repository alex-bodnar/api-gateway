package log

import (
	"fmt"
)

// InitLogger initializes common logger
func InitLogger(cfg Config, tags map[string]string) Logger {
	var mode, format, level int
	switch cfg.Mode {
	case "prod":
		mode = ModProd
	case "dev":
		mode = ModDev
	default:
		fmt.Println("Logger mod not set using dev mode")
		mode = ModDev
	}

	switch cfg.LogFormat {
	case "json":
		format = FormatJSON
	case "text":
		format = FormatConsole
	default:
		fmt.Println("Logger log format not set using text mode")
		format = FormatConsole
	}

	switch cfg.LogLevel {
	case "info":
		level = InfoLevel
	case "debug":
		level = DebugLevel
	case "error":
		level = ErrorLevel
	default:
		fmt.Println("Logger log level not set using debug level")
		level = DebugLevel
	}

	return NewEnv(
		mode,
		format,
		Tags(tags),
		Level(level),
		WithCaller(cfg.IncludeCallerMethod),
	)
}
