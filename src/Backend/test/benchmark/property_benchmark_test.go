package benchmark

import (
	"context"
	"testing"

	"github.com/Turgho/Aluguei/internal/infrastructure/persistence"
	"github.com/Turgho/Aluguei/test/testhelpers"
	"github.com/google/uuid"
)

func BenchmarkPropertyRepository_Create(b *testing.B) {
	db, err := testhelpers.SetupTestDB()
	if err != nil {
		b.Fatal(err)
	}

	repo := persistence.NewPropertyRepository(db)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		property := testhelpers.CreateTestProperty(uuid.New())
		err := repo.Create(ctx, property)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPropertyRepository_GetByID(b *testing.B) {
	db, err := testhelpers.SetupTestDB()
	if err != nil {
		b.Fatal(err)
	}

	repo := persistence.NewPropertyRepository(db)
	ctx := context.Background()

	// Setup test data
	property := testhelpers.CreateTestProperty(uuid.New())
	err = repo.Create(ctx, property)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := repo.GetByID(ctx, property.ID)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPropertyRepository_GetAll(b *testing.B) {
	db, err := testhelpers.SetupTestDB()
	if err != nil {
		b.Fatal(err)
	}

	repo := persistence.NewPropertyRepository(db)
	ctx := context.Background()

	// Setup test data
	ownerID := uuid.New()
	for i := 0; i < 100; i++ {
		property := testhelpers.CreateTestProperty(ownerID)
		err := repo.Create(ctx, property)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := repo.GetAll(ctx, 1, 20, "")
		if err != nil {
			b.Fatal(err)
		}
	}
}