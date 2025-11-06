package database

import (
	"strings"
	"time"

	"github.com/Turgho/Aluguei/internal/config"
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

var (
	dbInstance *DB
)

// Connect estabelece conexão com o PostgreSQL usando o logger Zap
func Connect(cfg *config.Config) (*DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	// Criar logger do GORM com Zap
	gormLogger := NewGormLogger(logger.Get().With(zap.String("component", "database")))

	// Ajustar nível de log baseado na configuração
	gormLogger.LogLevel = getGormLogLevel(cfg.LogLevel)
	if cfg.Environment == "development" {
		gormLogger.SlowThreshold = 200 * time.Millisecond
	} else {
		gormLogger.SlowThreshold = 500 * time.Millisecond
	}

	// Conectar ao banco
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger:                                   gormLogger,
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: false,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		logger.Error("Failed to connect to database",
			zap.String("database_url", maskDatabaseURL(cfg.DatabaseURL)),
			zap.Error(err),
		)
		return nil, err
	}

	// Configurar connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Failed to get SQL DB instance", zap.Error(err))
		return nil, err
	}

	dbConfig := cfg.DatabaseConfig()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Info("Database connection pool configured",
		zap.Int("max_idle_connections", dbConfig.MaxIdleConns),
		zap.Int("max_open_connections", dbConfig.MaxOpenConns),
	)

	// Executar migrations apenas em desenvolvimento
	if cfg.Environment == "development" {
		if err := runMigrations(db); err != nil {
			logger.Error("Failed to run database migrations", zap.Error(err))
			return nil, err
		}
		logger.Info("Database migrations completed successfully")
	}

	// Configurar callbacks para logging
	setupCallbacks(db)

	dbInstance = &DB{db}

	logger.Info("Database connected successfully",
		zap.String("dialect", "postgres"),
		zap.String("environment", cfg.Environment),
	)

	return dbInstance, nil
}

// getGormLogLevel converte nosso log level para o do GORM
func getGormLogLevel(level string) gormLogger.LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		return gormLogger.Info
	case "info":
		return gormLogger.Warn
	case "warn", "warning":
		return gormLogger.Warn
	case "error":
		return gormLogger.Error
	default:
		return gormLogger.Silent
	}
}

// maskDatabaseURL mascara a senha na URL do banco para logs
func maskDatabaseURL(url string) string {
	parts := strings.Split(url, "@")
	if len(parts) != 2 {
		return "***"
	}
	return "***@" + parts[1]
}

// runMigrations executa auto-migration
func runMigrations(db *gorm.DB) error {
	models := []interface{}{
		&models.User{},
		&models.Property{},
		&models.Contract{},
		&models.Payment{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

// setupCallbacks configura callbacks do GORM para logging adicional
func setupCallbacks(db *gorm.DB) {
	// Callback para log de criação
	db.Callback().Create().After("gorm:after_create").Register("logger:after_create", func(db *gorm.DB) {
		if db.Error == nil && db.Statement.Schema != nil {
			logger.Get().Debug("Database record created",
				zap.String("table", db.Statement.Schema.Table),
				zap.Int64("rows_affected", db.RowsAffected),
			)
		}
	})

	// Callback para log de updates
	db.Callback().Update().After("gorm:after_update").Register("logger:after_update", func(db *gorm.DB) {
		if db.Error == nil && db.Statement.Schema != nil {
			logger.Get().Debug("Database record updated",
				zap.String("table", db.Statement.Schema.Table),
				zap.Int64("rows_affected", db.RowsAffected),
			)
		}
	})

	// Callback para log de deletes
	db.Callback().Delete().After("gorm:after_delete").Register("logger:after_delete", func(db *gorm.DB) {
		if db.Error == nil && db.Statement.Schema != nil {
			logger.Get().Debug("Database record deleted",
				zap.String("table", db.Statement.Schema.Table),
				zap.Int64("rows_affected", db.RowsAffected),
			)
		}
	})
}

// GetDB retorna a instância do banco
func GetDB() *DB {
	return dbInstance
}

// Close fecha a conexão com o banco
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		logger.Error("Failed to get SQL DB for closing", zap.Error(err))
		return err
	}

	logger.Info("Closing database connection")
	return sqlDB.Close()
}

// HealthCheck verifica a saúde do banco
func (db *DB) HealthCheck() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		logger.Error("Failed to get SQL DB for health check", zap.Error(err))
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Error("Database health check failed", zap.Error(err))
		return err
	}

	stats := sqlDB.Stats()
	logger.Debug("Database connection pool stats",
		zap.Int("open_connections", stats.OpenConnections),
		zap.Int("in_use", stats.InUse),
		zap.Int("idle", stats.Idle),
	)

	return nil
}

// Transaction executa uma transação com logging
func (db *DB) Transaction(fc func(tx *gorm.DB) error) error {
	logger.Debug("Starting database transaction")
	start := time.Now()

	err := db.DB.Transaction(fc)

	elapsed := time.Since(start)
	if err != nil {
		logger.Error("Database transaction failed",
			zap.Duration("elapsed", elapsed),
			zap.Error(err),
		)
	} else {
		logger.Debug("Database transaction completed",
			zap.Duration("elapsed", elapsed),
		)
	}

	return err
}
