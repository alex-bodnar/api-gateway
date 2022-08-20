package log

type (
	// Logger common logger interface.
	Logger interface {
		Printf(string, ...interface{})
		// Info writes a information message.
		Info(...interface{})
		// Infof writes a formatted information message.
		Infof(string, ...interface{})
		// Infow writes a formatted information message with key,val pairs
		Infow(string, ...interface{})
		// Warning writes a warning message.
		Warning(...interface{})
		// Warningf writes a formatted warning message.
		Warningf(string, ...interface{})
		// Warningw writes a formatted information message with key,val pairs
		Warningw(string, ...interface{})
		// Error writes an error message.
		Error(...interface{})
		// Errorf writes a formatted error message.
		Errorf(string, ...interface{})
		// Errorw writes a formatted information message with key,val pairs
		Errorw(string, ...interface{})
		// Debug writes a debug message.
		Debug(...interface{})
		// Debugf writes a formatted debug message.
		Debugf(string, ...interface{})
		// Debugw writes a formatted information message with key,val pairs
		Debugw(string, ...interface{})
		// Fatal writes a fatal message.
		Fatal(...interface{})
		// Fatalf writes a formatted fatal message.
		Fatalf(string, ...interface{})
		// Fatalw writes a formatted information message with key,val pairs
		Fatalw(string, ...interface{})
		// Critical writes a critical message.
		Critical(v ...interface{})
		// Criticalf writes a formatted critical message.
		Criticalf(format string, v ...interface{})
		// Criticalw writes a formatted information message with key,val pairs
		Criticalw(s string, v ...interface{})
		// With add fields to be used for all logs
		With(f ...interface{}) Logger
	}
)
