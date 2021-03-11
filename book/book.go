package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/ketan-sonar/api-with-gofiber/database"
)

// Book ...
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks ...
func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

// GetBook ...
func GetBook(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}
	return c.JSON(book)
}

// NewBook ...
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	if book.Title == "" || book.Author == "" {
		return c.Status(500).SendString("Invalid format for creating a book")
	}
	db.Create(book)
	return c.JSON(*book)
}

// DeleteBook ...
func DeleteBook(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}
	db.Delete(&book)
	return c.SendString("Successfully deleted the book")
}
