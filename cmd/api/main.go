package main

import (
	echo "github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// @title           Service API
// @version         1.0
// @description     Production API service
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in              header
// @name            Authorization
func main() {
	var app = echo.New()

	app.Use(middleware.RequestLogger())
	app.Use(middleware.Recover())

	if err := app.Start(":8080"); err != nil {
		app.Logger.Error("failed to start server", "error", err)
	}
}
