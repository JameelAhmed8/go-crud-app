package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

func AuthorDeleteHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	authorId := qp.Get("author_id")
	Id, err := uuid.Parse(authorId)
	if err != nil {
		log.Printf("failed to parse authIdStr=%v: %v", Id, err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}
	err = models.DeleteAuthor(Id)
	if err != nil {
		log.Println("failed to delete author", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, nil)
}
