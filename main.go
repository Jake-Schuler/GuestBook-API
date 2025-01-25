package main

import (
	"time"

	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text   string `json:"Text"`
	Author string `json:"Author"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("gb-api.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	StartupTime := time.Now()

	db.AutoMigrate(&Message{})

	// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/messages", func(c echo.Context) error {
		var messages []Message
		db.Find(&messages)
		return c.JSON(200, messages)
	})

	e.GET("/totalMessages", func(c echo.Context) error {
		var count int64
		db.Model(&Message{}).Count(&count)
		return c.JSON(200, count)
	})

	e.POST("/newMessage", func(c echo.Context) error {
		var message Message
		if err := c.Bind(&message); err != nil {
			return c.JSON(400, echo.Map{"error": "Invalid JSON"})
		}
		if message.Text == "" || message.Author == "" {
			return c.JSON(400, echo.Map{"error": "Text and Author are required"})
		}
		if err := db.Create(&message).Error; err != nil {
			return c.JSON(500, echo.Map{"error": "Failed to create message"})
		}
		return c.JSON(201, echo.Map{"message": message})
	})

	e.GET("/uptime", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"uptime": time.Since(StartupTime).String()})
	})

	e.Static("/", "assets")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
