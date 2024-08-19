package main

import (
	"go-dice/environment"
	"go-dice/routes"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Static("/", "static")

	if environment.IsDevelopment() {
		e.Use(
			func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					if c.Request().URL.Path == "/css/style.css" {
						c.Response().Header().Set("Cache-Control", "no-store")
						c.Response().Header().Set("Pragma", "no-cache")
						c.Response().Header().Set("Expires", "0")
					}
					return next(c)
				}
			},
		)
	}
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${method} ${uri} ${status}
`,
	}))
	e.Use(middleware.Recover())

	routes.Init(e)

	e.Logger.Fatal(e.Start(":" + environment.EnvVar("PORT")))
}
