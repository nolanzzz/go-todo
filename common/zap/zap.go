package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
	// AddCaller if want to track trace
	logger = logger.WithOptions(zap.AddCaller())
	return logger
}

func getCore(filename string, level zapcore.LevelEnabler) zapcore.Core {
	writer := getWriter(filename)
	// Need three configs - Encoder, WriterSyncer, LogLevel
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder get zapcore.Encoder
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getWriter(filename string) zapcore.WriteSyncer {
	file, _ := os.Create(fmt.Sprintf("./log/%s", filename))
	return zapcore.AddSync(file)
}
