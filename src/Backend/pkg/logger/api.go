package logger

import "go.uber.org/zap"

// Get retorna o logger global
func Get() *zap.Logger {
	if globalLogger == nil {
		// Fallback para desenvolvimento
		development, _ := zap.NewDevelopment()
		return development
	}
	return globalLogger
}

// Sugar retorna o sugared logger
func Sugar() *zap.SugaredLogger {
	if sugarLogger == nil {
		development, _ := zap.NewDevelopment()
		return development.Sugar()
	}
	return sugarLogger
}

// Métodos helpers para uso rápido
func Debug(msg string, fields ...zap.Field) {
	Get().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Get().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Get().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Get().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Get().Fatal(msg, fields...)
}

// Métodos com contexto (sugared)
func Debugf(template string, args ...interface{}) {
	Sugar().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	Sugar().Errorf(template, args...)
}

// WithFields cria um logger com campos pré-definidos
func With(fields ...zap.Field) *zap.Logger {
	return Get().With(fields...)
}
