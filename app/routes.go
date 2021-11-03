package app

import (
	"echo-boilerplate/db"
	"echo-boilerplate/domain"
	"echo-boilerplate/service"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	dbClient := db.GetDbClient()
	helloRepositoryDb := domain.NewHelloRepositoryDb(dbClient)
	sh := SampleHandler{Service: service.DefaultHelloService{helloRepositoryDb}}
	e.GET("/", sh.PrintHello)
	return e
}
