package request

import (
	"Books/untils/structs"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bookscollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	bookscollection = client.Database("books").Collection("DataBooks")
}

func AddBooks(c *gin.Context) {
	name_books := c.PostForm("name_books")
	author := c.PostForm("author_books")
	price := c.PostForm("price")
	if name_books == "" && author == "" && price == "" {
		c.JSON(400, gin.H{
			"error": "Все поля должны быть заполнены",
		})
	}
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Некорекетный формат цены",
		})
	}
	books_data := structs.Books{
		Name:   name_books,
		Author: author,
		Price:  priceInt,
	}
	_, err = bookscollection.InsertOne(context.TODO(), books_data)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Книга была успешно добавлена",
	})
}
func FindBook(c *gin.Context) ([]structs.Books, error) {
	var DataBook []structs.Books
	var filter bson.D
	searchBy := c.PostForm("search_by")
	data_user := c.PostForm("search_value")
	switch searchBy {
	case "name":
		filter = bson.D{{"name", data_user}}
	case "author":
		filter = bson.D{{"author", data_user}}
	case "price":
		priceInt, err := strconv.Atoi(data_user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Неккоректный формат цены",
			})
		}
		filter = bson.D{{"price", priceInt}}
	}
	cursor, err := bookscollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Не удалось найти :", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &DataBook); err != nil {
		log.Println("Ошибка при чтении результатов поиска: ", err)
		return nil, err
	}
	return DataBook, nil
}
func BookOutput(c *gin.Context) []structs.Books {
	var books []structs.Books
	var filter bson.D
	filter = bson.D{{}}
	cursor, err := bookscollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Не удалось найти :", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &books); err != nil {
		log.Println("Ошибка при чтении результатов поиска: ", err)
		return nil
	}

	return books
}
func OutputDetail(c *gin.Context) structs.Books {
	var BooksDetail structs.Books
	var filter bson.D
	IdStr := c.Param("id")
	log.Println("ID из URL", IdStr)
	id, err := primitive.ObjectIDFromHex(IdStr)
	if err != nil {
		log.Println(nil)
	}
	filter = bson.D{{"_id", id}}
	err = bookscollection.FindOne(context.TODO(), filter).Decode(&BooksDetail)
	log.Println("Это выведено с booksdetail", BooksDetail)
	if err != nil {
		log.Println(err)
	}
	return BooksDetail
}
