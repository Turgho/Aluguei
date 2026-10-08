// internal/infra/repositories/user_repository_test.go
package repositories_test

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// newMockDB cria uma instância do GORM com banco mockado.
func newMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	require.NoError(t, err)

	return db, mock
}

// newFakeUser retorna um usuário fake para uso nos testes.
func newFakeUser() *entities.User {
	return &entities.User{
		ID:           uuid.New(),
		FirstName:    "João",
		LastName:     "Silva",
		CPF:          "000.000.000-00",
		Email:        "joao@email.com",
		Phone:        "11999999999",
		PasswordHash: "hash123",
		Role:         entities.RoleTenant,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func TestCreate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.New()))
	mock.ExpectCommit()

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	rows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "cpf", "email",
		"phone", "password_hash", "role", "created_at", "updated_at",
	}).AddRow(
		user.ID, user.FirstName, user.LastName, user.CPF, user.Email,
		user.Phone, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt,
	)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1`).
		WithArgs(user.ID.String(), 1).
		WillReturnRows(rows)

	result, err := repo.GetByID(user.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByEmail(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	rows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "cpf", "email",
		"phone", "password_hash", "role", "created_at", "updated_at",
	}).AddRow(
		user.ID, user.FirstName, user.LastName, user.CPF, user.Email,
		user.Phone, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt,
	)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1`).
		WithArgs(user.Email, 1).
		WillReturnRows(rows)

	result, err := repo.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByCPF(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	rows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "cpf", "email",
		"phone", "password_hash", "role", "created_at", "updated_at",
	}).AddRow(
		user.ID, user.FirstName, user.LastName, user.CPF, user.Email,
		user.Phone, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt,
	)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE cpf = \$1`).
		WithArgs(user.CPF, 1).
		WillReturnRows(rows)

	result, err := repo.GetByCPF(user.CPF)
	assert.NoError(t, err)
	assert.Equal(t, user.CPF, result.CPF)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Update(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "users" WHERE id = \$1`).
		WithArgs(user.ID.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Delete(user.ID.String())
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSearch(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	rows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "cpf", "email",
		"phone", "password_hash", "role", "created_at", "updated_at",
	}).AddRow(
		user.ID, user.FirstName, user.LastName, user.CPF, user.Email,
		user.Phone, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt,
	)

	mock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs("%joão%", "%joão%", "%joão%", "%joão%").
		WillReturnRows(rows)

	results, err := repo.Search("joão")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
