package main

import (
	"github.com/t01gyl0p/zolo-assignment-server/controllers"

	"github.com/t01gyl0p/zolo-assignment-server/types"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&types.Book{}, &types.Borrow{})

	bookController := controllers.BookController{DB: db}

	router.POST("/api/v1/booky/", bookController.AddBook)
	router.GET("/api/v1/booky/", bookController.BrowseBooks)
	router.DELETE("/api/v1/booky/:book_id/", bookController.DeleteBook)
	router.GET("/api/v1/borrow/", bookController.ListBorrows)
	router.POST("/api/v1/booky/:book_id/borrow", bookController.BorrowBook)
	router.PUT("/api/v1/booky/:book_id/borrow/:borrow_id", bookController.ReturnBook)

	router.Run(":8000")
}
