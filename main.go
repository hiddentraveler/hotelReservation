package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("mercii")
	app.Get("/", handleGreet)
	app.Listen(":5000")
}

func handleGreet(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Hello World"})
}
