package main

import (
	"fmt"
	"library_crud/routes"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	routes.GetRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error putting up the server. Error: %s", err.Error())
	}
}