package logger

type LoggerInterface interface {
	Debugf(format string, args ...interface{})
	Debug(string)

	Infof(format string, args ...interface{})
	Info(string)

	Printf(format string, args ...interface{})
	Print(string)

	Warnf(format string, args ...interface{})
	Warn(string)

	Errorf(format string, args ...interface{})
	Error(string)

	Fatalf(format string, args ...interface{})
	Fatal(string)

	Panicf(format string, args ...interface{})
	Panic(string)
}

type GooglyLoggerInterface interface {
	LoggerInterface
	WithData(map[string]interface{}) LoggerInterface
}
