//package main
//
//import (
//	"context"
//	"github.com/gorilla/mux"
//	"goproj/config"
//	"goproj/database"
//	"goproj/handlers"
//	"net/http"
//)
//
//func main() {
//	conf := config.GetConfig()
//	ctx := context.TODO()
//
//	db := database.ConnectDB(ctx, conf.Mongo)
//	collection := db.Collection(conf.Mongo.Collection)
//
//	client := &database.TodoClient{
//		Col: collection,
//		Ctx: ctx,
//	}
//
//	r := mux.NewRouter()
//
//	r.HandleFunc("/todos", handlers.SearchTodos(client)).Methods("GET")
//	r.HandleFunc("/todos/{id}", handlers.GetTodo(client)).Methods("GET")
//	r.HandleFunc("/todos", handlers.InsertTodo(client)).Methods("POST")
//	r.HandleFunc("/todos/{id}", handlers.UpdateTodo(client)).Methods("PATCH")
//	r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client)).Methods("DELETE")
//
//	http.ListenAndServe(":8080", r)
//}
//
//

	package main

	type Notebook struct {
		brand string
		price int
		mouse string
	}

	func NewNotebook(b string, p int, m string) *Notebook {
		return &Notebook{
			brand: b,
			price: p,
			mouse: m,
		}
	}

	func(n *Notebook) Add() {
		n.price++
	}

	func main() {
		NewNotebook("asus", 50, "domryhfsdf")

	}

