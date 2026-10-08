package database

import (
	"fmt"

	"github.com/Turgho/Aluguei/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB inicializa a conexão com bando de dados e retorna a instância de conexão do GORM
func NewDB(databaseURL string) (*gorm.DB, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("URL do banco vazio")
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar o banco: %w", err)
	}

	conn, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("falha ao pegar instância do banco: %w", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("falha ao receber resposta do banco: %w", err)
	}

	logger.Log.Info("banco connectado com sucesso")
	return db, nil
}
