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
			Input         string   `form:"input"`
			CurrentResult string   `form:"currentResult"`
			Selection     []string `form:"selection"`
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

		selectedDice := []models.DiceIdentifier{}

		for _, thing := range u.Selection {
			parts := strings.Split(thing, "@")

			leftValue, err1 := strconv.Atoi(parts[0])
			rightValue, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				panic("oops")
			}

			selectedDice = append(selectedDice, models.DiceIdentifier{
				RollValue:   leftValue,
				FromTheLeft: rightValue,
			})
		}

		rawOutput, err := engine.RollDice(u.CurrentResult + u.Input)
		if err != nil {
			return err
		}

		var strNumbers []string
		for _, num := range rawOutput {
			strNumbers = append(strNumbers, strconv.Itoa(num))
		}

		rollAmounts := []int{0, 0, 0, 0, 0, 0}
		for _, num := range rawOutput {
			rollAmounts[num-1]++
		}

		rollStats := models.RollStats{
			Rolls:         rawOutput,
			RollAmounts:   rollAmounts,
			PreviousInput: u.Input,
			CurrentRoll:   "+[" + strings.Join(strNumbers, ",") + "]",
			Selection:     selectedDice,
		}

		return views.Form(rollStats).Render(c.Request().Context(), c.Response().Writer)
	})
}
