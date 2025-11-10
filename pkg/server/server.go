package server

import (
	"e-commerce/config"
	"e-commerce/pkg/database"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() {
	config.LoadConfig()
	database.NewDB()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Print("server running at localhost:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
