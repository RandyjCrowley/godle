package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "/Users/randycrowley/Volumes/godle/wordle/wordle.go"
)

func main() {
    e := echo.New()


    e.GET("/",func(c echo.Context) error {
        return c.String(http.StatusOK,"Deez Nuts\n")
    })
    


    e.GET("/start",wordle.start)







    e.Logger.Fatal(e.Start(":1323"))
}