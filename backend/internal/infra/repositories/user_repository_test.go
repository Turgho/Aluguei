// internal/infra/repositories/user_repository_test.go
package repositories_test

import (
	"database/sql/driver"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	domain "github.com/Turgho/Aluguei/internal/domain/repositories"
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
	now := time.Now()

	return &entities.User{
		ID:            uuid.New(),
		FirstName:     "João",
		LastName:      "Silva",
		CPF:           "000.000.000-00",
		Email:         "joao@email.com",
		Phone:         "11999999999",
		PasswordHash:  "hash123",
		Role:          entities.RoleTenant,
		IsActive:      true,
		EmailVerified: false,
		CreatedAt:     now,
		UpdatedAt:     now,
		DeletedAt:     gorm.DeletedAt{},
	}
}

// userColumns são as colunas padrão retornadas nas queries de SELECT.
var userColumns = []string{
	"id", "first_name", "last_name", "cpf", "email",
	"phone", "password_hash", "role", "is_active", "email_verified",
	"created_at", "updated_at", "deleted_at",
}

// userRow converte um usuário em uma linha compatível com sqlmock.
func userRow(u *entities.User) []driver.Value {
	return []driver.Value{
		u.ID,
		u.FirstName,
		u.LastName,
		u.CPF,
		u.Email,
		u.Phone,
		u.PasswordHash,
		u.Role,
		u.IsActive,
		u.EmailVerified,
		u.CreatedAt,
		u.UpdatedAt,
		nil,
	}
}

// ── Leitura por chave única ───────────────────────────────────────────────────

func TestGetByID(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	rows := sqlmock.NewRows(userColumns).AddRow(userRow(user)...)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1 AND "users"\."deleted_at" IS NULL LIMIT \$2`).
		WithArgs(user.ID, 1).
		WillReturnRows(rows)

	result, err := repo.GetByID(user.ID)

	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByEmail(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	rows := sqlmock.NewRows(userColumns).AddRow(userRow(user)...)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"\."deleted_at" IS NULL LIMIT \$2`).
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

	rows := sqlmock.NewRows(userColumns).AddRow(userRow(user)...)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE cpf = \$1 AND "users"\."deleted_at" IS NULL LIMIT \$2`).
		WithArgs(user.CPF, 1).
		WillReturnRows(rows)

	result, err := repo.GetByCPF(user.CPF)

	assert.NoError(t, err)
	assert.Equal(t, user.CPF, result.CPF)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// ── Escrita ───────────────────────────────────────────────────────────────────

func TestCreate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()

	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(
			user.FirstName,
			user.LastName,
			user.CPF,
			user.Email,
			user.Phone,
			user.PasswordHash,
			user.Role,
			user.IsActive,
			user.EmailVerified,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			user.ID,
		).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "last_login_at"}).
				AddRow(user.ID, nil),
		)

	mock.ExpectCommit()

	err := repo.Create(user)

	assert.NoError(t, err)
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
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "deleted_at"=\$1 WHERE id = \$2 AND "users"\."deleted_at" IS NULL`).
		WithArgs(sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.Delete(user.ID)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// ── Busca com filtros ─────────────────────────────────────────────────────────

func TestSearch(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	// COUNT
	countRow := sqlmock.NewRows([]string{"count"}).AddRow(1)

	mock.ExpectQuery(`SELECT count\(\*\) FROM "users"`).
		WillReturnRows(countRow)

	// SELECT
	rows := sqlmock.NewRows(userColumns).AddRow(userRow(user)...)

	mock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(
			"%joão%",
			"%joão%",
			"%joão%",
			"%joão%",
			20,
		).
		WillReturnRows(rows)

	results, total, err := repo.Search(
		domain.UserFilters{
			Search:   "joão",
			Page:     1,
			PageSize: 20,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, results, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSearchWithRoleFilter(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	role := entities.RoleTenant

	countRow := sqlmock.NewRows([]string{"count"}).AddRow(1)

	mock.ExpectQuery(`SELECT count\(\*\) FROM "users"`).
		WithArgs(role).
		WillReturnRows(countRow)

	rows := sqlmock.NewRows(userColumns).AddRow(userRow(user)...)

	mock.ExpectQuery(`SELECT \* FROM "users"`).
		WithArgs(role, 20).
		WillReturnRows(rows)

	results, total, err := repo.Search(
		domain.UserFilters{
			Role:     &role,
			Page:     1,
			PageSize: 20,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, results, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// ── Operações de conta ────────────────────────────────────────────────────────

func TestSetActive(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "is_active"=\$1,"updated_at"=\$2 WHERE id = \$3`).
		WithArgs(false, sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.SetActive(user.ID, false)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSetEmailVerified(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "email_verified"=\$1,"updated_at"=\$2 WHERE id = \$3`).
		WithArgs(true, sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.SetEmailVerified(user.ID)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateLastLogin(t *testing.T) {
	db, mock := newMockDB(t)
	repo := repositories.NewUserRepository(db)
	user := newFakeUser()
	now := time.Now()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "last_login_at"=\$1,"updated_at"=\$2 WHERE id = \$3`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.UpdateLastLogin(user.ID, now)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
