package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type CategoriesListResp struct {
	Category []*CategoryListRespObj `json:"categories"`
}

type CategoryListRespObj struct {
	ID   uuid.UUID `json:"category_id"`
	Name string    `json:"name"`
}

func CategoriesListHandler(w http.ResponseWriter, r *http.Request) {

	tups, err := models.ListCategories()
	if err != nil {
		log.Println("failed to list categories", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// Convert to endpoint output format
	data := make([]*CategoryListRespObj, 0, len(tups))
	for _, tup := range tups {
		data = append(data, &CategoryListRespObj{
			ID:   tup.CategoryID,
			Name: tup.Name,
		})
	}
	utils.WriteJSONResponse(w, http.StatusOK, data)
}
