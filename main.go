package main

import (
	"fmt"
	"todoList/db"
	"todoList/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("HELLO Wrld")
}
func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=54321"
	db.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection failed!")
	}
	fmt.Println("connected succesfully!")
	db.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("migrated db!")
}
func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodo)

}
func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}
