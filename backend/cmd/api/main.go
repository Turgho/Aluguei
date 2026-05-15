// Package main é o ponto de entrada da aplicação Aluguei.
package main

// main inicializa e executa o servidor HTTP da aplicação.
//
// A inicialização completa das dependências (banco de dados, logger,
// roteador e rotas) é delegada para [NewServer].
func main() {
	server := NewServer()
	server.Run()
}
