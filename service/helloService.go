package service

import (
	"echo-boilerplate/domain"
	"echo-boilerplate/dto"
)

type HelloService interface {
	NewMessage(req dto.HelloRequest) (*dto.HelloResponse, error)
}

type DefaultHelloService struct {
	repo domain.HelloRepository
}

func (s DefaultHelloService) NewMessage(req dto.HelloRequest) (*dto.HelloResponse, error) {
	var res dto.HelloResponse
	res.Message = "Hello " + req.Name
	return &res, nil
}

func NewHelloService() HelloService {
	return &DefaultHelloService{}
}
