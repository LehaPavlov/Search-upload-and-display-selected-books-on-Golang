package main

import (
	"Books/render"
	"Books/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", render.RenderMainPage)
	router.GET("/book/:id", render.BookDetail)
	router.GET("/add_books", render.RenderAddBooks)
	router.POST("/add_books", request.AddBooks)
	router.GET("/find_book", render.FindBook)
	router.POST("/find_book", func(c *gin.Context) {
		books, err := request.FindBook(c)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "findbook.html", books)
	})
	router.Run()
}
