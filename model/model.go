package book

import (
	"database/sql"
	"fmt"
	"library_crud/db"
)

type Book struct {
	Book_id  int
	Title    string
	Author   string
	Quantity int
}

func GetAllBooks() []Book {
	db := db.PostgresConnection()

	allBooks, err := db.Query("SELECT * FROM public.book ORDER BY book_id ASC")

	if err != nil {
		panic(err.Error())
	}

	book := Book{}
	bookList := []Book{}

	for allBooks.Next() {
		var book_id, quantity int
		var title, author string

		err = allBooks.Scan(&book_id, &title, &author, &quantity)

		if err != nil {
			panic(err.Error())
		}

		book.Book_id = book_id
		book.Title = title
		book.Author = author
		book.Quantity = quantity

		bookList = append(bookList, book)

	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	return bookList
}

func CreateBook(title, author string, quantity int) {
	db := db.PostgresConnection()

	createBook, err := db.Prepare("INSERT INTO public.book(title, author, quantity)	VALUES ($1, $2, $3);")

	if err != nil {
		fmt.Printf("Error preparing query. %s", err.Error())
		return
	}

	createBook.Exec(title, author, quantity)

	db.Close()
}

func DeleteBook(id int) {
	db := db.PostgresConnection()

	deleteBook, err := db.Prepare("DELETE FROM public.book WHERE book_id = $1")

	if err != nil {
		fmt.Printf("Error preparing query. %s", err.Error())
		return
	}

	deleteBook.Exec(id)

	db.Close()
}

func GetBookById(id int) Book {
	db := db.PostgresConnection()

	book, err := db.Query("SELECT * FROM public.book WHERE book_id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	editBook := Book{}

	for book.Next() {
		var id, quantity int
		var title, author string

		err = book.Scan(&id, &title, &author, &quantity)

		if err != nil {
			fmt.Println("Error scaning book atributes")
			panic(err.Error())
		}

		editBook.Book_id = id
		editBook.Title = title
		editBook.Author = author
		editBook.Quantity = quantity

	}

	defer db.Close()

	return editBook
}

func UpdateBook(id, quantity int, title, author string) {
	db := db.PostgresConnection()

	bookUpdate, err := db.Prepare("update public.book set title=$1, author=$2, quantity=$3 where book_id=$4")

	if err != nil {
		fmt.Printf("Error preparing query. %s", err.Error())
		return
	}

	bookUpdate.Exec(title, author, quantity, id)

	defer db.Close()
}