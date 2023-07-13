package routes

import (
	"github.com/gorilla/mux"
	author "github.com/JameelAhmed8/go-crud-app/pkg/controllers/author"
)

func SetRoutes(router *mux.Router) {
	// Author Routes
	authorRouter := router.PathPrefix("/api/author/v1").Subrouter()
	authorRouter.HandleFunc("/create", author.AuthorCreateHandler).Methods("POST")
	authorRouter.HandleFunc("/list", author.AuthorsListHandler).Methods("GET")
	authorRouter.HandleFunc("/get", author.AuthorGetHandler).Methods("GET")
	authorRouter.HandleFunc("/update", author.AuthorUpdateHandler).Methods("PUT")
	authorRouter.HandleFunc("/delete", author.AuthorDeleteHandler).Methods("DELETE")
}
