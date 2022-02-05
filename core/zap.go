package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"todo/global"
)

func Zap() *zap.Logger {
	// Create log directory if not exist already
	if _, err := os.Stat("./log"); err != nil {
		_ = os.Mkdir("./log", os.ModePerm)
	}
	// Debug Level
	debugPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.DebugLevel
	})
	// Info Level
	infoPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.InfoLevel
	})
	// Error Level
	errorPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	cores := [...]zapcore.Core{
		getCore("debug.log", debugPriority),
		getCore("info.log", infoPriority),
		getCore("error.log", errorPriority),
	}
	logger := zap.New(zapcore.NewTee(cores[:]...))
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getCore(filename string, level zapcore.LevelEnabler) zapcore.Core {
	writer := getWriter(filename)
	// Need three configs - Encoder, WriterSyncer, LogLevel
	return zapcore.NewCore(getEncoder(), writer, level)
}

func getEncoder() zapcore.Encoder {
	var config zapcore.EncoderConfig
	if global.CONFIG.Zap.Format == "json" {
		config = zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

func getWriter(filename string) zapcore.WriteSyncer {
	file, _ := os.Create(fmt.Sprintf("./log/%s", filename))
	return zapcore.AddSync(file)
}
