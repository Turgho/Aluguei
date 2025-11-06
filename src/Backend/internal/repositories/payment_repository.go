package repositories

import (
	"time"

	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) *errors.AppError
	FindByID(id string) (*models.Payment, *errors.AppError)
	FindByContractID(contractID string) ([]models.Payment, *errors.AppError)
	FindOverdue() ([]models.Payment, *errors.AppError)
	FindByPeriod(startDate, endDate time.Time) ([]models.Payment, *errors.AppError)
	Update(payment *models.Payment) *errors.AppError
	Delete(id string) *errors.AppError
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

func (r *paymentRepository) FindByContractID(contractID string) ([]models.Payment, *errors.AppError) {
	var payments []models.Payment
	err := r.db.Where("contract_id = ?", contractID).
		Preload("Contract").
		Order("due_date DESC").
		Find(&payments).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar pagamentos por contrato", err)
	}
	return payments, nil
}

func (r *paymentRepository) FindOverdue() ([]models.Payment, *errors.AppError) {
	var payments []models.Payment
	err := r.db.Where("status = ? AND due_date < ?",
		models.PaymentStatusOverdue, time.Now().Format("2006-01-02")).
		Preload("Contract").
		Preload("Contract.Property").
		Preload("Contract.Tenant").
		Find(&payments).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar pagamentos em atraso", err)
	}
	return payments, nil
}

func (r *paymentRepository) FindByPeriod(startDate, endDate time.Time) ([]models.Payment, *errors.AppError) {
	var payments []models.Payment
	err := r.db.Where("due_date BETWEEN ? AND ?", startDate, endDate).
		Preload("Contract").
		Preload("Contract.Property").
		Preload("Contract.Tenant").
		Order("due_date ASC").
		Find(&payments).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar pagamentos por período", err)
	}
	return payments, nil
}

// Sobrescrever FindByID para incluir preloads
func (r *paymentRepository) FindByID(id string) (*models.Payment, *errors.AppError) {
	var payment models.Payment
	err := r.db.Preload("Contract").
		Preload("Contract.Property").
		Preload("Contract.Tenant").
		First(&payment, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("pagamento", id)
		}
		return nil, errors.NewDatabaseError("erro ao buscar pagamento por ID", err)
	}
	return &payment, nil
}

// Sobrescrever FindAll para incluir preloads
func (r *paymentRepository) FindAll() ([]models.Payment, *errors.AppError) {
	var payments []models.Payment
	err := r.db.Preload("Contract").
		Preload("Contract.Property").
		Preload("Contract.Tenant").
		Find(&payments).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar todos os pagamentos", err)
	}
	return payments, nil
}
