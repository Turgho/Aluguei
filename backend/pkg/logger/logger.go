// Package logger fornece uma instância global de logger estruturado
// baseada em [go.uber.org/zap], com comportamento adaptado ao ambiente
// de execução da aplicação.
package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log é a instância global do logger, inicializada por [Init].
// Deve ser chamada somente após [Init] ter sido executado.
var Log *zap.Logger

// Init inicializa a instância global [Log] de acordo com o ambiente informado.
//
// Ambientes suportados:
//
//   - "production": logger JSON estruturado, nível Info, timestamp ISO8601.
//     Adequado para ingestão por ferramentas como Datadog, Loki ou CloudWatch.
//
//   - qualquer outro valor (ex: "development"): logger colorido no terminal,
//     nível Debug, com stacktrace automático em erros.
//     Adequado para desenvolvimento local.
func Init(env string) {
	var encoderCfg zapcore.EncoderConfig

	if env == "production" {
		// Produção: JSON estruturado para facilitar ingestão por ferramentas de observabilidade.
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
		// Desenvolvimento: saída colorida e legível no terminal com stacktrace em erros.
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
