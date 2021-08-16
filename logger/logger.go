package logger

// LoggerInterface is a sub interface implemented by GooglyLoggerInterface
type LoggerInterface interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Warnf(format string, args ...interface{})
	Warn(args ...interface{})

	Errorf(format string, args ...interface{})
	Error(args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})

	Panicf(format string, args ...interface{})
	Panic(args ...interface{})
}

// GooglyLoggerInterface is the logger interface implemented in all services for logging
//
// Any logger that implements this interface becomes a GooglyLogger.
//
// If the logger of your choice that doesnt implement this interface,
// you have to write a simple binding
//
// WithData can be used to log extra data along with log string.
// it adds extra info to your logs
type GooglyLoggerInterface interface {
	LoggerInterface
	WithData(map[string]interface{}) LoggerInterface
}
