// Package main demonstra o uso do Options Pattern em Go.
// O Options Pattern é uma técnica para passar configurações opcionais
// a uma função ou construtor de forma elegante e extensível.
package main

import "fmt"

// Options é um tipo funcional que representa uma opção de configuração.
// Cada opção recebe um ponteiro para Server e modifica seus campos.
type Options func(*Server)

// Server representa um servidor com configurações obrigatórias e opcionais.
// maxConn, id e tls são parâmetros obrigatórios passados diretamente no construtor.
// optionA e optionB são parâmetros opcionais configurados via Options.
type Server struct {
	maxConn int
	id      string
	tls     bool
	optionA bool
	optionB bool
}

// WithOptionA retorna uma Options que define o valor de optionA no Server.
// Convenção: funções de opção seguem o prefixo "With" seguido do nome do campo.
func WithOptionA(b bool) Options {
	return func(s *Server) {
		s.optionA = b
	}
}

// WithOptionB retorna uma Options que define o valor de optionB no Server.
func WithOptionB(b bool) Options {
	return func(s *Server) {
		s.optionB = b
	}
}

// NewServer é o construtor do Server. Recebe parâmetros obrigatórios (maxConn, id, tls)
// e um variadic de Options para configurações adicionais.
// Isso permite criar o Server com apenas os parâmetros essenciais e,
// opcionalmente, customizá-lo sem alterar a assinatura da função.
func NewServer(maxConn int, id string, tls bool, opts ...Options) *Server {
	// Inicializa o Server com valores padrão para os campos opcionais.
	s := &Server{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
		optionA: false,
		optionB: false,
	}
	// Aplica cada opção fornecida sobre o Server criado.
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {

	s1 := NewServer(100, "server1", true)
	fmt.Println("With default values:", s1)

	s2 := NewServer(100, "server1", true, WithOptionA(true), WithOptionB(true))
	fmt.Println("With options enabled:", s2)

}
