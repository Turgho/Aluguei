package utils_test

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateCPF_ValidCPF(t *testing.T) {
	// CPF válido para teste
	validCPF := "11144477735"
	result := utils.ValidateCPF(validCPF)
	assert.True(t, result)
}

func TestValidateCPF_InvalidCPF(t *testing.T) {
	// CPF inválido (todos os dígitos iguais)
	invalidCPF := "11111111111"
	result := utils.ValidateCPF(invalidCPF)
	assert.False(t, result)
}

func TestValidateCPF_InvalidLength(t *testing.T) {
	// CPF com tamanho inválido
	shortCPF := "123456789"
	result := utils.ValidateCPF(shortCPF)
	assert.False(t, result)
}

func TestValidateCPF_WithMask(t *testing.T) {
	// CPF com máscara
	maskedCPF := "111.444.777-35"
	result := utils.ValidateCPF(maskedCPF)
	assert.True(t, result)
}