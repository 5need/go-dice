package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	ProductRoutes(e)
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
}
