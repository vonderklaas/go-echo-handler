package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World")

	server := echo.New()

	server.Use(MiddlewareAdmin)

	server.GET("/status", Handler)

	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {

	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	s := fmt.Sprintf("Amount of days: %d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

func MiddlewareAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("User-Role")

		if val == "admin" {
			log.Println("Red Button User Detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
