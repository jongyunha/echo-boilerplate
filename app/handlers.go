package app

import (
	"echo-boilerplate/dto"
	"echo-boilerplate/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SampleHandler struct {
	Service service.HelloService
}

// @Summary Print Hello
// @Description Print Hello
// @Accept  string
// @Param   name     path    string     true        "name"
// @Success 200 {object} dto.HelloResponse
// @Router  / [get]
func (sh SampleHandler) PrintHello(c echo.Context) error {
	var res dto.HelloRequest
	name := c.Param("name")
	res.Name = name
	err := res.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := sh.Service.NewMessage(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
