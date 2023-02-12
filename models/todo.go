package models

import (
	"todoList/db"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm: "primarykey" json:"id"`
	Title     string `json: "title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := db.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "erreur", "message": "Todo not found!", "data": err})
	}
	return c.JSON(&todo)
}

func CreateTodo(c *fiber.Ctx) error {
	db := db.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "erreur", "message": "invalid input", "data": err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "erreur", "message": "cannot create todo", "data": err})
	}
	return c.JSON(&todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	type UpdatedTodo struct {
		Title     string `json: "title"`
		Completed bool   `json:"completed"`
	}
	id := c.Params("id")
	db := db.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "erreur", "message": "Todo not found!", "data": err})
	}
	var updatedTodo UpdatedTodo
	err = c.BodyParser(&updatedTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "input Error", "message": "ERREUR!", "data": err})
	}
	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	db.Save(&todo)
	return c.JSON(&todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "erreur", "message": "Todo not found!", "data": err})
	}
	db.Delete(&todo)
	return c.SendStatus(200)
}
