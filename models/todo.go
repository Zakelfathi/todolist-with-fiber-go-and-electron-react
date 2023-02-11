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
