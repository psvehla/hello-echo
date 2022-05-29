package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/psvehla/hello-echo/models"
)

// HelloHello - hello hello
func (c *Container) HelloHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello, " + ctx.Param("name") + ".",
	})
}

// HelloOpenapiJson - Download ./gen/http/openapi.json
func (c *Container) HelloOpenapiJson(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
