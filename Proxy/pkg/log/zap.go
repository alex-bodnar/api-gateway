package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

const (
	ModDev = iota
	ModProd
)

const (
	FormatJSON = iota
	FormatConsole
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

// Option options for zap setup
type Option func(z *zap.Logger) *zap.Logger

// Level set log level
func Level(logLevel int) Option {
	return func(z *zap.Logger) *zap.Logger {
		return z.WithOptions(zap.IncreaseLevel(zapcore.Level(logLevel)))
	}
}

// Name sets logger name
func Name(name string) Option {
	return func(z *zap.Logger) *zap.Logger {
		return z.Named(name)
	}
}

// WithCaller enables output of caller method
func WithCaller(enabled bool) Option {
	return func(z *zap.Logger) *zap.Logger {
		return z.WithOptions(zap.WithCaller(enabled))
	}
}

// Tags adds default tags for logger
func Tags(fields map[string]string) Option {
	return func(z *zap.Logger) *zap.Logger {
		if len(fields) > 0 {
			zapFields := make([]zap.Field, 0, len(fields))
			for k, v := range fields {
				zapFields = append(zapFields, zap.String(k, v))
			}
			return z.With(zapFields...)
		}
		return z
	}
}

// Output sets output
func Output(o io.Writer) Option {
	return func(z *zap.Logger) *zap.Logger {
		if o != nil {
			return z.WithOptions(zap.ErrorOutput(zapcore.AddSync(o)))
		}
		return z
	}
}

type zapLog struct {
	log *zap.SugaredLogger
}

// NewWithOutput creates logger depending on mode
func NewWithOutput(mode, format int, infoOut io.Writer, errOut io.Writer, options ...Option) Logger {
	var cfg zapcore.EncoderConfig
	var enc zapcore.Encoder

	switch mode {
	case ModDev:
		cfg = zap.NewDevelopmentEncoderConfig()
	case ModProd:
		cfg = zap.NewProductionEncoderConfig()
	default:
		cfg = zap.NewDevelopmentEncoderConfig()
	}

	switch format {
	case FormatJSON:
		enc = zapcore.NewJSONEncoder(cfg)
	case FormatConsole:
		enc = zapcore.NewConsoleEncoder(cfg)
	default:
		enc = zapcore.NewConsoleEncoder(cfg)
	}

	coreJson := zapcore.NewCore(
		enc,
		zapcore.AddSync(infoOut),
		zap.DebugLevel,
	)
	logger := zap.New(coreJson)
	logger = logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.ErrorOutput(zapcore.AddSync(errOut)))

	if len(options) > 0 {
		for _, o := range options {
			logger = o(logger)
		}
	}

	return zapLog{log: logger.Sugar()}
}

// NewEnv creates logger depending on mode
func NewEnv(mode, format int, options ...Option) Logger {
	var cfg zapcore.EncoderConfig
	var enc zapcore.Encoder

	switch mode {
	case ModDev:
		cfg = zap.NewDevelopmentEncoderConfig()
	case ModProd:
		cfg = zap.NewProductionEncoderConfig()
	default:
		cfg = zap.NewDevelopmentEncoderConfig()
	}

	cfg.FunctionKey = "F"

	switch format {
	case FormatJSON:
		enc = zapcore.NewJSONEncoder(cfg)
	case FormatConsole:
		enc = zapcore.NewConsoleEncoder(cfg)
	default:
		enc = zapcore.NewConsoleEncoder(cfg)
	}

	coreJson := zapcore.NewCore(
		enc,
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)
	logger := zap.New(coreJson)
	logger = logger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.ErrorOutput(zapcore.AddSync(os.Stderr)),
		zap.AddStacktrace(zap.FatalLevel),
	)

	if len(options) > 0 {
		for _, o := range options {
			logger = o(logger)
		}
	}

	return zapLog{log: logger.Sugar()}
}

// New creates new logger instance
func New(options ...Option) Logger {
	var logger *zap.Logger

	logger, _ = zap.NewDevelopment()
	if len(options) > 0 {
		for _, o := range options {
			logger = o(logger)
		}
	}

	logger = logger.WithOptions(zap.AddCallerSkip(1))

	return zapLog{log: logger.Sugar()}
}

func NewTest(t zaptest.TestingT, opt ...zaptest.LoggerOption) Logger {
	return zapLog{log: zaptest.NewLogger(t, opt...).Sugar()}
}

func (z zapLog) Printf(s string, v ...interface{}) {
	z.log.Infof(s, v...)
}

func (z zapLog) Info(v ...interface{}) {
	z.log.Info(v...)
}

func (z zapLog) Infof(s string, v ...interface{}) {
	z.log.Infof(s, v...)
}

func (z zapLog) Infow(s string, v ...interface{}) {
	z.log.Infow(s, v...)
}

func (z zapLog) Warning(v ...interface{}) {
	z.log.Warn(v...)
}

func (z zapLog) Warningw(s string, v ...interface{}) {
	z.log.Warnw(s, v...)
}

func (z zapLog) Warningf(s string, v ...interface{}) {
	z.log.Warnf(s, v...)
}

func (z zapLog) Critical(v ...interface{}) {
	z.log.Fatal(v...)
}

func (z zapLog) Criticalf(s string, v ...interface{}) {
	z.log.Fatalf(s, v...)
}

func (z zapLog) Criticalw(s string, v ...interface{}) {
	z.log.Fatalw(s, v...)
}

func (z zapLog) Error(v ...interface{}) {
	z.log.Error(v...)
}

func (z zapLog) Errorf(s string, v ...interface{}) {
	z.log.Errorf(s, v...)
}

func (z zapLog) Errorw(s string, v ...interface{}) {
	z.log.Errorw(s, v...)
}

func (z zapLog) Debug(v ...interface{}) {
	z.log.Debug(v...)
}

func (z zapLog) Debugf(s string, v ...interface{}) {
	z.log.Debugf(s, v...)
}

func (z zapLog) Debugw(s string, v ...interface{}) {
	z.log.Debugw(s, v...)
}

func (z zapLog) Fatal(v ...interface{}) {
	z.log.Fatal(v...)
}

func (z zapLog) Fatalw(s string, v ...interface{}) {
	z.log.Fatalw(s, v...)
}

func (z zapLog) Fatalf(s string, v ...interface{}) {
	z.log.Fatalf(s, v...)
}

func (z zapLog) With(fields ...interface{}) Logger {
	return zapLog{log: z.log.With(fields)}
}

func (z zapLog) Flush() error {
	return z.log.Sync()
}
