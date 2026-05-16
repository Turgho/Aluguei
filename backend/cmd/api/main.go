// Package main é o ponto de entrada da aplicação Aluguei.
package main

// @title           Aluguei API
// @version         0.3.1
// @description     API do sistema Aluguei.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	server := NewServer()
	server.Run()
}
