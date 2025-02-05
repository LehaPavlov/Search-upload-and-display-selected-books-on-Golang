Application Description:

I have developed a web application for managing a book library using Go, Gin (a web framework) and MongoDB (a database).

Application Functionality:

Viewing the list of books:
The main page displays a list of all the books stored in the MongoDB database.
For each book, the title, author and price are displayed.
Adding books:
There is a page for adding new books.
The user can enter the title, author and price of the book.
After submitting the form, the book is saved in the MongoDB database.
Searching for books:
The user can search for a book by author.
Book Details:
For each book in the list, there is a “More details” link.
Clicking on this link opens a separate page with detailed information about the book: title, author, price and a unique identifier (ID) of the book in the database.
Technologies and Components:

Go: The programming language in which the application is written.
Gin: A lightweight web framework for Go that makes it easy to build web applications.
MongoDB: A NoSQL database used to store book data.
HTML Templates: Used to create the HTML pages that the user sees.
BSON: A binary data format used by MongoDB to store documents.
go.mongodb.org/mongo-driver/bson/primitive: A package for working with the type used to represent an identifier in MongoDB.primitive.ObjectID_id
