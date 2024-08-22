package wordle

import (
	"github.com/labstack/echo/v4"
)

func Start(c echo.Context) error {
	// Function implementation
	return c.String(http.StatusOK, "Hello from GetMyData!")
}
