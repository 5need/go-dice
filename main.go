package main

import (
	"fmt"
	"go-dice/engine"
	"go-dice/routes"
	"net/http"
	"os"

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
	args := os.Args

	if len(args) > 1 {
		for _, arg := range args[1:] {
			fmt.Println(engine.RollDice(arg))
		}
		return
	}

	e := echo.New()
	e.Static("/", "static")

	if env := os.Getenv("ENV"); env != "production" {
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

	routes.Routes(e)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if code == http.StatusNotFound {
			c.String(http.StatusNotFound, "404 Not Found")
			return
		}
		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Logger.Fatal(e.Start(":3000"))
}
