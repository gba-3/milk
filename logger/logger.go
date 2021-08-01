package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

var loggerConf = zap.Config{
	OutputPaths: []string{"/var/log/ap/ap.log"},
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

func SetupLogger(level string) {
	l, err := NewLogger(level)
	if err != nil {
		log.Println(err)
	}
	Log = l
}

func NewLogger(level string) (*zap.Logger, error) {
	conf := loggerConf
	conf.Level = zap.NewAtomicLevelAt(convertLevel(level))
	return conf.Build()
}
