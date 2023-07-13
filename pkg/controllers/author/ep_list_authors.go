package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type AuthorListResp struct {
	Authors []*AuthorListRespObj `json:"authors"`
}

type AuthorListRespObj struct {
	AuthorID    uuid.UUID `json:"author_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Affiliation string    `json:"affiliation"`
	Biography   string    `json:"biography"`
	Website     string    `json:"website"`
}

func AuthorsListHandler(w http.ResponseWriter, r *http.Request) {

	tups, err := models.ListAuthors()
	if err != nil {
		log.Println("failed to list authors", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// Convert to endpoint output format
	data := make([]*AuthorListRespObj, 0, len(tups))
	for _, tup := range tups {
		data = append(data, &AuthorListRespObj{
			AuthorID:    tup.AuthorID,
			Name:        tup.Name,
			Email:       tup.Email,
			Affiliation: tup.Affiliation,
			Biography:   tup.Biography,
			Website:     tup.Website,
		})
	}
	utils.WriteJSONResponse(w, http.StatusOK, data)
}
