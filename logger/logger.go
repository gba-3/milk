package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log = NewLogger()

func NewLogger() *zap.Logger {
	logConfig := zap.Config{
		OutputPaths: []string{"./ap.log"},
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, err := logConfig.Build()
	if err != nil {
		return zap.NewExample()
	}
	if err := log.Sync(); err != nil {
		return zap.NewExample()
	}
	return log
}
