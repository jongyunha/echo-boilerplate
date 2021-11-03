package dto

import "errors"

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func (r HelloRequest) Validate() error {
	if HelloRequest.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
