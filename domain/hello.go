package domain

type Hello struct {
	Name string
}

type HelloRepository interface {
	FindByName(name string) (*Hello, error)
}
