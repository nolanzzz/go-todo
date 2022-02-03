package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() *zap.Logger {
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
	return zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
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
