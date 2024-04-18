package log

import (
	"io"

	"git.virjar.com/Junhiee/anilismei/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// init conig

type ZapEncConfig struct {
}

func NewGinZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func NewGinZapwriteSyncer() io.Writer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/zap.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return lumberJackLogger

}

func InitLogger() {
	enc := zapcore.NewJSONEncoder(NewGinZapEncoderConfig())
	ws := zapcore.AddSync(NewGinZapwriteSyncer())
	lv := zapcore.DebugLevel
	core := zapcore.NewCore(enc, ws, lv)
	global.ZLOG = zap.New(core)
}
