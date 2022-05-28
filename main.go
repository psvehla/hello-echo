package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/psvehla/hello-echo/handlers"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// HelloHello - hello hello
	e.GET("/hello/:name", c.HelloHello)

	// HelloOpenapiJson - Download ./gen/http/openapi.json
	e.GET("/openapi.json", c.HelloOpenapiJson)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
