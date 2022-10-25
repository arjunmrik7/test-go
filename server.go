package main

import (


	"github.com/gofiber/fiber/v2"
	
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func initDb() {
// 	var err error
// 	database.DBconn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect to db")
// 	}

// 	fmt.Println("db connection successfully opened")
// }

func main() {

	app := fiber.New(fiber.Config{
		// Prefork:       true,
		// CaseSensitive: true,
		// StrictRouting: true,
		// ServerHeader:  "Fiber",
		// AppName:       "ShareAuto Go v1.0.1",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

	app.Listen(":3000")

	//initDb()
}
