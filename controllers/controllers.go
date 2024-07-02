package controllers

import (
	"fmt"
	book "library_crud/model"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	library := book.GetAllBooks()

	if err := temp.ExecuteTemplate(w, "Home", library); err != nil {
		fmt.Printf("Error executing home template. Error: %s", err.Error())
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		quantity := r.FormValue("quantity")

		quantityConverted, err := strconv.Atoi(quantity)

		if err != nil {
			fmt.Println("Error converting qunatity")
			w.WriteHeader(http.StatusBadRequest)
		}

		book.CreateBook(title, author, quantityConverted)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	
	if err := temp.ExecuteTemplate(w, "Create", nil); err != nil {
		fmt.Printf("Error executing insert template. %s", err.Error())
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	bookId := r.URL.Query().Get("id")

	bookIdConverted, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Printf("Error converting book id to int. Error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	book.DeleteBook(bookIdConverted)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func EditBook(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "POST" {
		bookId := r.URL.Query().Get("id")
		title := r.FormValue("title")
		author := r.FormValue("author")
		quantity := r.FormValue("quantity")

		convertedBookId, err := strconv.Atoi(bookId)
	
		if err != nil {
			fmt.Printf("Error converting book id to int. Error: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			fmt.Printf("Error converting book id to int. Error: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}

		book.UpdateBook(convertedBookId, convertedQuantity, title, author)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	
	bookId := r.URL.Query().Get("id")

	convertedBookId , err := strconv.Atoi(bookId)
	if err != nil {
		panic(err.Error())
	}

	book := book.GetBookById(convertedBookId)

	if err = temp.ExecuteTemplate(w, "Edit", book); err != nil {
		fmt.Printf("Error executing edit template. Error: %s", err.Error())
	}
}