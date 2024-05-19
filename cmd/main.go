package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/Allyedge/paste/config"
	"github.com/Allyedge/paste/db"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	cfg := config.NewConfig()

	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer database.Close()

	err = db.Migrate(database)
	if err != nil {
		panic(err)
	}

	hr := cfg.InitializeHandlers(cfg.InitializeRepositories(database))

	app.GET("/", hr.HomeHandler.HandleHomeShow)
	app.POST("/", hr.HomeHandler.HandleHomeCreate)
	app.GET("/:id", hr.CodeHandler.HandleCodeShow)

	app.Start(":4000")
}
