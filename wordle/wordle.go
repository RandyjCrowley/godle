package wordle

import (
//	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type WordleResponse struct {
	ID  string `json:"id"`
	Word string `json:"word"`
}

func Start(c echo.Context) error {
	id := uuid.Must(uuid.NewRandom()).String()
	word := get_word()
	response := WordleResponse{ID: id, Word: word}
	return c.JSON(http.StatusOK, response)
}
