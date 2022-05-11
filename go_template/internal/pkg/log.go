package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	zap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILog interface {
	Info(message string)
	Error(message string)
	Close()
}

type Log struct {
}

var logInst *Log
var logger *zap.Logger
var zap_config zap.Config

func LogInit(config IConfig, args *IArgs) {
	logInst = &Log{}

	config_path := config.ConfigurePath()

	bytes, _ := os.ReadFile(config_path)

	println(config_path)

	var cfg zap.Config
	err = json.Unmarshal(bytes, &cfg)

	if err != nil {
		panic(err)
	}

	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05")))
	}
	cfg.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	zap_config = cfg
	logger, err = cfg.Build(zap.WrapCore(zapCore))
}

func zapCore(c zapcore.Core) zapcore.Core {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   zap_config.OutputPaths[1],
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     100, //days
		Compress:   true,
	})

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap_config.EncoderConfig),
		w,
		zap.DebugLevel,
	)
	cores := zapcore.NewTee(c, core)

	return cores
}

func LogInst() ILog {
	return logInst
}

func (l *Log) Info(message string) {
	logger.Info(message)
}

func (l *Log) Error(message string) {
	logger.Error(message)
}

func (l *Log) Close() {
	//nothing
}
