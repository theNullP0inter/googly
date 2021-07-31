package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger() *logrus.Logger {
	logger := logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: true},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}
	dsn := viper.Get("SENTRY_DSN")
	if dsn == nil {
		return &logger
	}

	if dsn != nil {
		hook, err := logrus_sentry.NewSentryHook(dsn.(string), []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		timeout := viper.GetInt("SENTRY_TIMEOUT")
		hook.Timeout = time.Duration(timeout) * time.Second
		hook.StacktraceConfiguration.Enable = true

		if err == nil {
			logger.Hooks.Add(hook)
		}
	}

	return &logger
}

func NewSentryClient() *sentry.Client {
	dsn := viper.Get("SENTRY_DSN")
	if dsn == nil {
		return nil
	}

	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:   dsn.(string),
		Debug: true,
	})
	if err != nil {
		fmt.Println("Fatal")
		fmt.Println(err)
	}
	return client
}
