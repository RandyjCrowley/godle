package wordle

import (
	//	"encoding/json"
	"net/http"
	"os"
    "strings"
    "bufio"

    "fmt"
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
    AddWord(word)
	response := WordleResponse{ID: id, Word: word}
	return c.JSON(http.StatusOK, response)
}

func GetFile() (*os.File) {
    file, err := os.OpenFile("word.list", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return file
}

func CheckWord(word string) bool {
    file := GetFile()
    
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), word) {
			return true
		}
	}
    return false
}

func AddWord(word string) bool {
    file := GetFile()
    if (! CheckWord(word)) {
        _, err := file.WriteString(word)
        if err != nil {
            fmt.Println("Error writing to file:", err)
        }
        defer file.Close()
        return true
    }
    defer file.Close()
    return false
}

