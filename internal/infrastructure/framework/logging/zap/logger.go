package zap

import (
	"app/internal/core/port/logging"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewZapLogger() *zap.SugaredLogger {
	var z *zap.Logger
	var err error
	devMode := os.Getenv("ZAP_DEV_MODE") == "true"

	if devMode {
		z, err = zap.NewDevelopment()
	} else {
		z, err = zap.NewProduction()
	}

	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}

	return z.Sugar()
}

func ConfigureZap(zap *zap.SugaredLogger, g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(ginzap.Ginzap(zap.Desugar(), time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(zap.Desugar(), true))
}

func NewLogger(zap *zap.SugaredLogger) logging.Logger {
	return &Logger{
		zap: zap,
	}
}

type Logger struct {
	zap *zap.SugaredLogger
}

func (l *Logger) Debug(args ...interface{}) {
	l.zap.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.zap.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.zap.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.zap.Error(args...)
}
