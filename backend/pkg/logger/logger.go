package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(env string) {
	var encoderCfg zapcore.EncoderConfig

	// Log via JSON estruturado em PRODUÇÃO
	// Log colorido legível no terminal em DESENVOLVIMENTO
	if env == "production" {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		)

		Log = zap.New(core, zap.AddCaller())
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zap.DebugLevel,
		)
		Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	}
}
