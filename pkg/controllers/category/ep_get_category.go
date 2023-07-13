package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type CategoryGetResp struct {
	Id   uuid.UUID `json:"category_id"`
	Name string    `json:"name"`
}

func CategoryGetHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	catId := qp.Get("category_id")
	Id, err := uuid.Parse(catId)
	if err != nil {
		log.Printf("failed to parse authIdStr=%v: %v", Id, err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}
	tup, err := models.GetCategoryByID(Id)
	if err != nil {
		log.Println("failed to get category", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// Convert to endpoint output format
	data := CategoryGetResp{
		Name: tup.Name,
		Id:   tup.CategoryID,
	}

	utils.WriteJSONResponse(w, http.StatusOK, data)

}
