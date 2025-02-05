package render

import (
	"Books/request"
	"Books/untils/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderMainPage(c *gin.Context) {
	var DataBooks []structs.Books
	DataBooks = request.BookOutput(c)
	log.Println(DataBooks)
	c.HTML(http.StatusOK, "index.html", DataBooks)
}

func RenderAddBooks(c *gin.Context) {
	c.HTML(http.StatusOK, "add_books.html", nil)
}

func FindBook(c *gin.Context) {
	c.HTML(http.StatusOK, "findbook.html", nil)
}

func BookDetail(c *gin.Context) {
	var BookID structs.Books
	BookID = request.OutputDetail(c)
	c.HTML(http.StatusOK, "bookdetail.html", BookID)

}
