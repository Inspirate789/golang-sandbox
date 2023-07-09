package interfaces

import "github.com/Inspirate789/golang-sandbox/internal/models"

type Client interface {
	Open(target string) error
	Close() error
	models.CalculatorExternal
}

type Server interface {
	Serve() error
	Close() error
}
