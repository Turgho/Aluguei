package repositories

import (
	"strings"

	"github.com/Turgho/Aluguei/internal/errors"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) *errors.AppError {
	err := r.db.Create(entity).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "Duplicate entry") {
			return errors.NewAlreadyExistsError("registro", "chave única", "")
		}
		return errors.NewDatabaseError("erro ao criar registro", err)
	}
	return nil
}

func (r *BaseRepository[T]) FindByID(id string) (*T, *errors.AppError) {
	var entity T
	err := r.db.First(&entity, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("registro", id)
		}
		return nil, errors.NewDatabaseError("erro ao buscar registro por ID", err)
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindByIDWithPreloads(id string, preloads ...string) (*T, *errors.AppError) {
	var entity T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.First(&entity, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("registro", id)
		}
		return nil, errors.NewDatabaseError("erro ao buscar registro por ID com preloads", err)
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, *errors.AppError) {
	var entities []T
	err := r.db.Find(&entities).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar todos os registros", err)
	}
	return entities, nil
}

func (r *BaseRepository[T]) FindAllWithPreloads(preloads ...string) ([]T, *errors.AppError) {
	var entities []T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.Find(&entities).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar todos os registros com preloads", err)
	}
	return entities, nil
}

func (r *BaseRepository[T]) Update(entity *T) *errors.AppError {
	err := r.db.Save(entity).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "Duplicate entry") {
			return errors.NewAlreadyExistsError("registro", "chave única", "")
		}
		return errors.NewDatabaseError("erro ao atualizar registro", err)
	}
	return nil
}

func (r *BaseRepository[T]) Delete(id string) *errors.AppError {
	err := r.db.Delete(new(T), "id = ?", id).Error
	if err != nil {
		return errors.NewDatabaseError("erro ao deletar registro", err)
	}
	return nil
}

// Método auxiliar para queries customizadas
func (r *BaseRepository[T]) Where(query interface{}, args ...interface{}) *gorm.DB {
	return r.db.Where(query, args...)
}
