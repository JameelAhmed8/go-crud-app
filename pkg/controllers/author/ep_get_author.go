package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type AuthorGetResp struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Affiliation string `json:"affiliation"`
	Biography   string `json:"biography"`
	Website     string `json:"website"`
}

func AuthorGetHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	authorId := qp.Get("author_id")
	Id, err := uuid.Parse(authorId)
	if err != nil {
		log.Printf("failed to parse authIdStr=%v: %v", Id, err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}
	tup, err := models.GetAuthor(Id)
	if err != nil {
		log.Println("failed to get athor", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// Convert to endpoint output format
	data := AuthorGetResp{
		Name:        tup.Name,
		Email:       tup.Email,
		Affiliation: tup.Affiliation,
		Biography:   tup.Biography,
		Website:     tup.Website,
	}

	utils.WriteJSONResponse(w, http.StatusOK, data)

}
