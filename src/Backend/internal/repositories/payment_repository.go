package repositories

import (
	"time"

	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	FindByID(id string) (*models.Payment, error)
	FindByContractID(contractID string) ([]models.Payment, error)
	FindOverdue() ([]models.Payment, error)
	FindByPeriod(startDate, endDate time.Time) ([]models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id string) error
}

type paymentRepository struct {
	*BaseRepository[models.Payment]
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		BaseRepository: NewBaseRepository[models.Payment](db),
		db:             db,
	}
}

func (r *paymentRepository) FindByContractID(contractID string) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("contract_id = ?", contractID).
		Preload("Contract").
		Order("due_date DESC").
		Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) FindOverdue() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("status = ? AND due_date < ?",
		models.PaymentStatusOverdue, time.Now().Format("2006-01-02")).
		Preload("Contract").
		Preload("Contract.Property").
		Preload("Contract.Tenant").
		Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) FindByPeriod(startDate, endDate time.Time) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("due_date BETWEEN ? AND ?", startDate, endDate).
		Preload("Contract").
		Preload("Contract.Property").
		Preload("Contract.Tenant").
		Order("due_date ASC").
		Find(&payments).Error
	return payments, err
}
