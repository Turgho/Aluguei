package seeds

import (
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) SeedAll() error {
	if err := s.SeedOwners(); err != nil {
		return err
	}
	if err := s.SeedTenants(); err != nil {
		return err
	}
	if err := s.SeedProperties(); err != nil {
		return err
	}
	if err := s.SeedContracts(); err != nil {
		return err
	}
	if err := s.SeedPayments(); err != nil {
		return err
	}
	return nil
}

func (s *Seeder) SeedOwners() error {
	// Check if owners already exist
	var count int64
	s.db.Model(&entities.Owner{}).Count(&count)
	if count > 0 {
		return nil
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	birthDate1 := time.Date(1980, 5, 15, 0, 0, 0, 0, time.UTC)
	birthDate2 := time.Date(1975, 8, 22, 0, 0, 0, 0, time.UTC)

	owners := []*entities.Owner{
		{
			ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
			Name:      "João Silva",
			Email:     "joao.silva@email.com",
			Password:  string(hashedPassword),
			Phone:     "+5511999999999",
			CPF:       "12345678901",
			BirthDate: &birthDate1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440002"),
			Name:      "Maria Santos",
			Email:     "maria.santos@email.com",
			Password:  string(hashedPassword),
			Phone:     "+5511888888888",
			CPF:       "98765432100",
			BirthDate: &birthDate2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440003"),
			Name:      "Carlos Oliveira",
			Email:     "carlos.oliveira@email.com",
			Password:  string(hashedPassword),
			Phone:     "+5511777777777",
			CPF:       "11122233344",
			BirthDate: nil,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return s.db.Create(&owners).Error
}

func (s *Seeder) SeedTenants() error {
	var count int64
	s.db.Model(&entities.Tenant{}).Count(&count)
	if count > 0 {
		return nil
	}

	birthDate1 := time.Date(1990, 3, 10, 0, 0, 0, 0, time.UTC)
	birthDate2 := time.Date(1985, 12, 5, 0, 0, 0, 0, time.UTC)

	tenants := []*entities.Tenant{
		{
			ID:        uuid.MustParse("660e8400-e29b-41d4-a716-446655440001"),
			Name:      "Ana Costa",
			Email:     "ana.costa@email.com",
			Phone:     "+5511666666666",
			CPF:       "55566677788",
			BirthDate: &birthDate1,
			OwnerID:   uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.MustParse("660e8400-e29b-41d4-a716-446655440002"),
			Name:      "Pedro Almeida",
			Email:     "pedro.almeida@email.com",
			Phone:     "+5511555555555",
			CPF:       "99988877766",
			BirthDate: &birthDate2,
			OwnerID:   uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.MustParse("660e8400-e29b-41d4-a716-446655440003"),
			Name:      "Lucia Ferreira",
			Email:     "lucia.ferreira@email.com",
			Phone:     "+5511444444444",
			CPF:       "33344455566",
			BirthDate: nil,
			OwnerID:   uuid.MustParse("550e8400-e29b-41d4-a716-446655440002"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return s.db.Create(&tenants).Error
}

func (s *Seeder) SeedProperties() error {
	var count int64
	s.db.Model(&entities.Property{}).Count(&count)
	if count > 0 {
		return nil
	}

	properties := []*entities.Property{
		{
			ID:          uuid.MustParse("770e8400-e29b-41d4-a716-446655440001"),
			OwnerID:     uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
			Title:       "Apartamento Centro - 2 Quartos",
			Description: "Apartamento bem localizado no centro da cidade com 2 quartos e 1 banheiro",
			Address:     "Rua das Flores, 123",
			City:        "São Paulo",
			State:       "SP",
			ZipCode:     "01234567",
			Bedrooms:    2,
			Bathrooms:   1,
			Area:        65,
			RentAmount:  1800.00,
			Status:      entities.PropertyStatusAvailable,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.MustParse("770e8400-e29b-41d4-a716-446655440002"),
			OwnerID:     uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
			Title:       "Casa Vila Madalena - 3 Quartos",
			Description: "Casa espaçosa na Vila Madalena com quintal e garagem",
			Address:     "Rua Harmonia, 456",
			City:        "São Paulo",
			State:       "SP",
			ZipCode:     "05435000",
			Bedrooms:    3,
			Bathrooms:   2,
			Area:        120,
			RentAmount:  3200.00,
			Status:      entities.PropertyStatusRented,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.MustParse("770e8400-e29b-41d4-a716-446655440003"),
			OwnerID:     uuid.MustParse("550e8400-e29b-41d4-a716-446655440002"),
			Title:       "Kitnet Liberdade",
			Description: "Kitnet compacta e funcional no bairro da Liberdade",
			Address:     "Rua da Glória, 789",
			City:        "São Paulo",
			State:       "SP",
			ZipCode:     "01510001",
			Bedrooms:    1,
			Bathrooms:   1,
			Area:        25,
			RentAmount:  950.00,
			Status:      entities.PropertyStatusAvailable,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.MustParse("770e8400-e29b-41d4-a716-446655440004"),
			OwnerID:     uuid.MustParse("550e8400-e29b-41d4-a716-446655440003"),
			Title:       "Cobertura Jardins - 4 Quartos",
			Description: "Cobertura de luxo nos Jardins com terraço e piscina",
			Address:     "Av. Paulista, 1000",
			City:        "São Paulo",
			State:       "SP",
			ZipCode:     "01310100",
			Bedrooms:    4,
			Bathrooms:   3,
			Area:        200,
			RentAmount:  8500.00,
			Status:      entities.PropertyStatusMaintenance,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return s.db.Create(&properties).Error
}

func (s *Seeder) SeedContracts() error {
	var count int64
	s.db.Model(&entities.Contract{}).Count(&count)
	if count > 0 {
		return nil
	}

	contracts := []*entities.Contract{
		{
			ID:            uuid.MustParse("880e8400-e29b-41d4-a716-446655440001"),
			PropertyID:    uuid.MustParse("770e8400-e29b-41d4-a716-446655440002"),
			TenantID:      uuid.MustParse("660e8400-e29b-41d4-a716-446655440001"),
			StartDate:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			MonthlyRent:   3200.00,
			PaymentDueDay: 5,
			Status:        entities.ContractStatusActive,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            uuid.MustParse("880e8400-e29b-41d4-a716-446655440002"),
			PropertyID:    uuid.MustParse("770e8400-e29b-41d4-a716-446655440003"),
			TenantID:      uuid.MustParse("660e8400-e29b-41d4-a716-446655440003"),
			StartDate:     time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			EndDate:       time.Date(2024, 5, 31, 0, 0, 0, 0, time.UTC),
			MonthlyRent:   950.00,
			PaymentDueDay: 10,
			Status:        entities.ContractStatusExpired,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	return s.db.Create(&contracts).Error
}

func (s *Seeder) SeedPayments() error {
	var count int64
	s.db.Model(&entities.Payment{}).Count(&count)
	if count > 0 {
		return nil
	}

	paidAmount1 := 3200.00
	paidDate1 := time.Date(2024, 1, 3, 10, 30, 0, 0, time.UTC)
	paidAmount2 := 3200.00
	paidDate2 := time.Date(2024, 2, 4, 14, 15, 0, 0, time.UTC)

	payments := []*entities.Payment{
		{
			ID:         uuid.MustParse("990e8400-e29b-41d4-a716-446655440001"),
			ContractID: uuid.MustParse("880e8400-e29b-41d4-a716-446655440001"),
			Amount:     3200.00,
			DueDate:    time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
			PaidAmount: &paidAmount1,
			PaidDate:   &paidDate1,
			Status:     entities.PaymentStatusPaid,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.MustParse("990e8400-e29b-41d4-a716-446655440002"),
			ContractID: uuid.MustParse("880e8400-e29b-41d4-a716-446655440001"),
			Amount:     3200.00,
			DueDate:    time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			PaidAmount: &paidAmount2,
			PaidDate:   &paidDate2,
			Status:     entities.PaymentStatusPaid,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.MustParse("990e8400-e29b-41d4-a716-446655440003"),
			ContractID: uuid.MustParse("880e8400-e29b-41d4-a716-446655440001"),
			Amount:     3200.00,
			DueDate:    time.Date(2024, 3, 5, 0, 0, 0, 0, time.UTC),
			PaidAmount: nil,
			PaidDate:   nil,
			Status:     entities.PaymentStatusPending,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.MustParse("990e8400-e29b-41d4-a716-446655440004"),
			ContractID: uuid.MustParse("880e8400-e29b-41d4-a716-446655440001"),
			Amount:     3200.00,
			DueDate:    time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
			PaidAmount: nil,
			PaidDate:   nil,
			Status:     entities.PaymentStatusOverdue,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	return s.db.Create(&payments).Error
}