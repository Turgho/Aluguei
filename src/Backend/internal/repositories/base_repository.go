package repositories

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) FindByID(id string) (*T, error) {
	var entity T
	err := r.db.First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindByIDWithPreloads(id string, preloads ...string) (*T, error) {
	var entity T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) FindAllWithPreloads(preloads ...string) ([]T, error) {
	var entities []T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id string) error {
	return r.db.Delete(new(T), "id = ?", id).Error
}
