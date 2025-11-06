package database

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// GormLogger implementa a interface de logger do GORM usando o Zap
type GormLogger struct {
	zapLogger        *zap.Logger
	LogLevel         logger.LogLevel
	SlowThreshold    time.Duration
	SkipCallerOffset int
}

// NewGormLogger cria uma nova instância do logger do GORM
func NewGormLogger(zapLogger *zap.Logger) *GormLogger {
	return &GormLogger{
		zapLogger:        zapLogger,
		LogLevel:         logger.Warn,
		SlowThreshold:    100 * time.Millisecond,
		SkipCallerOffset: 3, // Ajuste para pular as chamadas do GORM interno
	}
}

// LogMode define o nível de log
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs informações
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.zapLogger.Info(msg,
			zap.String("type", "gorm_info"),
			zap.Any("data", data),
			zap.String("caller", l.getCaller()),
		)
	}
}

// Warn logs avisos
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.zapLogger.Warn(msg,
			zap.String("type", "gorm_warn"),
			zap.Any("data", data),
			zap.String("caller", l.getCaller()),
		)
	}
}

// Error logs erros
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.zapLogger.Error(msg,
			zap.String("type", "gorm_error"),
			zap.Any("data", data),
			zap.String("caller", l.getCaller()),
		)
	}
}

// Trace logs queries SQL
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	// Campos comuns
	fields := []zap.Field{
		zap.String("type", "gorm_sql"),
		zap.Duration("elapsed", elapsed),
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.String("caller", l.getCaller()),
	}

	switch {
	case err != nil && l.LogLevel >= logger.Error:
		l.zapLogger.Error("SQL Error", append(fields, zap.Error(err))...)

	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		l.zapLogger.Warn("Slow SQL Query", append(fields,
			zap.String("slow_threshold", l.SlowThreshold.String()),
		)...)

	case l.LogLevel >= logger.Info:
		l.zapLogger.Debug("SQL Query", fields...)
	}
}

// getCaller retorna informações do caller para debugging
func (l *GormLogger) getCaller() string {
	// Pular as chamadas internas do GORM para encontrar o caller real
	for i := l.SkipCallerOffset; i < l.SkipCallerOffset+5; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.Contains(file, "gorm.io/gorm") && !strings.Contains(file, "database/gorm_logger.go")) {
			return fmt.Sprintf("%s:%d", filepath.Base(file), line)
		}
	}
	return "unknown"
}
