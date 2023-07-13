package routes

import (
	article "github.com/JameelAhmed8/go-crud-app/pkg/controllers/article"
	author "github.com/JameelAhmed8/go-crud-app/pkg/controllers/author"
	category "github.com/JameelAhmed8/go-crud-app/pkg/controllers/category"

	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	// Author Routes
	authorRouter := router.PathPrefix("/api/author/v1").Subrouter()
	authorRouter.HandleFunc("/create", author.AuthorCreateHandler).Methods("POST")
	authorRouter.HandleFunc("/list", author.AuthorsListHandler).Methods("GET")
	authorRouter.HandleFunc("/get", author.AuthorGetHandler).Methods("GET")
	authorRouter.HandleFunc("/update", author.AuthorUpdateHandler).Methods("PUT")
	authorRouter.HandleFunc("/delete", author.AuthorDeleteHandler).Methods("DELETE")

	// // Article Routes
	articleRouter := router.PathPrefix("/api/article/v1").Subrouter()
	articleRouter.HandleFunc("/create", article.ArticleCreateHandler).Methods("POST")
	articleRouter.HandleFunc("/list", article.ArticlesListHandler).Methods("GET")
	articleRouter.HandleFunc("/get", article.ArticleGetHandler).Methods("GET")
	articleRouter.HandleFunc("/update", article.ArticleUpdateHandler).Methods("PUT")
	articleRouter.HandleFunc("/delete", article.ArticleDeleteHandler).Methods("DELETE")

	// Category Routes
	categoryRouter := router.PathPrefix("/api/category/v1").Subrouter()
	categoryRouter.HandleFunc("/create", category.CategoryCreateHandler).Methods("POST")
	categoryRouter.HandleFunc("/list", category.CategoriesListHandler).Methods("GET")
	categoryRouter.HandleFunc("/get", category.CategoryGetHandler).Methods("GET")
	categoryRouter.HandleFunc("/update", category.CategoryUpdateHandler).Methods("PUT")
	categoryRouter.HandleFunc("/delete", category.CategoryDeleteHandler).Methods("DELETE")
}
