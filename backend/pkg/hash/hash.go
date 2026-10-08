// Package hash fornece utilitários para hashing e verificação de senhas
// utilizando o algoritmo Argon2id, recomendado pelo OWASP para armazenamento
// seguro de senhas.
package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

// params define os parâmetros de custo do Argon2id.
// Valores baseados nas recomendações do OWASP.
type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var defaultParams = &params{
	memory:      64 * 1024, // 64MB
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

// HashPassword gera um hash seguro da senha utilizando Argon2id.
//
// O hash retornado contém o salt e os parâmetros embutidos, no formato:
//
//	$argon2id$v=19$m=65536,t=3,p=2$<salt>$<hash>
//
// Exemplo:
//
//	hash, err := hash.HashPassword("minha-senha")
func HashPassword(password string) (string, error) {
	salt := make([]byte, defaultParams.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("erro ao gerar salt: %w", err)
	}

	h := argon2.IDKey(
		[]byte(password),
		salt,
		defaultParams.iterations,
		defaultParams.memory,
		defaultParams.parallelism,
		defaultParams.keyLength,
	)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(h)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultParams.memory,
		defaultParams.iterations,
		defaultParams.parallelism,
		encodedSalt,
		encodedHash,
	), nil
}

// VerifyPassword verifica se a senha corresponde ao hash armazenado.
//
// Retorna true se a senha for válida, false caso contrário.
//
// Exemplo:
//
//	match, err := hash.VerifyPassword("minha-senha", hashArmazenado)
func VerifyPassword(password, encodedHash string) (bool, error) {
	p, salt, h, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	candidate := argon2.IDKey(
		[]byte(password),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength,
	)

	// Comparação em tempo constante para evitar timing attacks
	if subtle.ConstantTimeCompare(h, candidate) == 1 {
		return true, nil
	}

	return false, nil
}

// decodeHash extrai os parâmetros, salt e hash de um hash Argon2id codificado.
func decodeHash(encodedHash string) (*params, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, fmt.Errorf("hash inválido")
	}

	// Esperado:
	// ["", "argon2id", "v=19", "m=65536,t=3,p=2", "<salt>", "<hash>"]
	if parts[1] != "argon2id" {
		return nil, nil, nil, fmt.Errorf("algoritmo inválido: %s", parts[1])
	}

	if !strings.HasPrefix(parts[2], "v=") {
		return nil, nil, nil, fmt.Errorf("versão inválida")
	}

	version, err := strconv.Atoi(strings.TrimPrefix(parts[2], "v="))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("erro ao ler versão: %w", err)
	}
	if version != argon2.Version {
		return nil, nil, nil, fmt.Errorf("versão argon2 não suportada: %d", version)
	}

	p := &params{}
	var mem, iter, par int

	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &iter, &par)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("erro ao ler parâmetros: %w", err)
	}

	p.memory = uint32(mem)
	p.iterations = uint32(iter)
	p.parallelism = uint8(par)

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("erro ao decodificar salt: %w", err)
	}

	h, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("erro ao decodificar hash: %w", err)
	}

	p.keyLength = uint32(len(h))
	p.saltLength = uint32(len(salt))

	return p, salt, h, nil
}
