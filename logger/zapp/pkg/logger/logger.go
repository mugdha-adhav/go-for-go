package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logFile = "temp.log"
)

var Log *zap.SugaredLogger

func New() error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Define log location
	consoleDebugging := zapcore.Lock(os.Stdout)
	fileDebugging := zapcore.Lock(f)

	// Create custom zap config
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(setUserConfigs()), consoleDebugging, zap.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), fileDebugging, zap.DebugLevel),
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

	defer Log.Sync()
	return nil
}

func setUserConfigs() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey: "message",
	}
}
