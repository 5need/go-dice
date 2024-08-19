package routes

import (
	"fmt"
	"go-dice/engine"
	"go-dice/models"
	"go-dice/views"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return views.Main().Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("/", func(c echo.Context) error {
		type PostRequest struct {
			Input         string `form:"input"`
			CurrentResult string `form:"currentResult"`
		}

		u := new(PostRequest)

		if err := c.Bind(u); err != nil {
			fmt.Println("failed to bind")
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(u); err != nil {
			fmt.Println("failed to validate")
			return c.NoContent(http.StatusBadRequest)
		}

		rawOutput := engine.RollDice(u.CurrentResult + u.Input)

		var strNumbers []string
		for _, num := range rawOutput {
			strNumbers = append(strNumbers, strconv.Itoa(num))
		}

		rollAmounts := []int{0, 0, 0, 0, 0, 0}
		for _, num := range rawOutput {
			rollAmounts[num-1]++
		}

		rollStats := models.RollStats{
			Rolls:       rawOutput,
			RollAmounts: rollAmounts,
		}

		views.Form(u.Input, "["+strings.Join(strNumbers, ",")+"]").Render(c.Request().Context(), c.Response().Writer)
		return views.DiceBox(rollStats).Render(c.Request().Context(), c.Response().Writer)
	})
}
