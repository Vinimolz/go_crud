package routes

import (
	"net/http"
	"library_crud/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/create", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.EditBook)
}