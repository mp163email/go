package hello

import "fmt"

type Service interface {
	Greet(name string) string
	Ping() string
}

type service struct {
}

func NewHelloService() Service {
	return &service{}
}

func (s *service) Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello %s", name)
}

func (s *service) Ping() string {
	return "pong"
}
