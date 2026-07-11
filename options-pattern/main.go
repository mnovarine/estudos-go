package main

import "fmt"

type Options func(*Server)

type Server struct {
	maxConn int
	id      string
	tls     bool
	optionA bool
	optionB bool
}

func WithOptionA(b bool) Options {
	return func(s *Server) {
		s.optionA = b
	}
}

func WithOptionB(b bool) Options {
	return func(s *Server) {
		s.optionB = b
	}
}

func NewServer(maxConn int, id string, tls bool, opts ...Options) *Server {
	s := &Server{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
		optionA: false,
		optionB: false,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {

	s := NewServer(100, "server1", true)
	fmt.Println("s:", s)

}
