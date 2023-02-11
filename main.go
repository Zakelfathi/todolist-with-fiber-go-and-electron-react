package main

import (
	"github.com/gofiber/fiber/v2"
)
func helloWorld(c *fiber.Ctx) error{
	return c.sendString("HELLO Wrld")
}
func main(){
	app := fiber.New()
	app.Get("/", helloWorld)
	app.Listen(":8000")
}