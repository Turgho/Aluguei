package entities_test

import "testing"

// assertNoError valida se nenhum erro foi retornado.
func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("esperava nil error, recebeu %v", err)
	}
}

// assertError valida se um erro foi retornado.
func assertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("esperava erro, recebeu nil")
	}
}
