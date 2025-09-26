package main

import (
	"log"
	"net/http"
	"todo-Api/internal/handler"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	http.HandleFunc("/add", handler.PostHandler)
	http.HandleFunc("/list", handler.GetHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)

	log.Println("Started server on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
