package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/JameelAhmed8/go-crud-app/pkg/routes"
)

const (
	serverPort = "8080"
)

func main() {
	r := mux.NewRouter()
	registerRoutes(r)

	addr := ":" + serverPort
	log.Printf("Server listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))

}

func registerRoutes(r *mux.Router) {
	routes.SetRoutes(r)
}
