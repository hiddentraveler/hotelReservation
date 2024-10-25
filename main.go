package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "Listen address of api server")
	app := fiber.New()
	fmt.Println("mercii")
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", handleUsers)
	app.Get("/", handleGreet)
	app.Listen(*listenAddr)
}

func handleUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Neo"})
}

func handleGreet(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Hello World"})
}
