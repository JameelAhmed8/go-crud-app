package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

func CategoryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	catId := qp.Get("category_id")
	fmt.Println(catId)
	Id, err := uuid.Parse(catId)
	if err != nil {
		log.Printf("failed to parse catStr=%v: %v", Id, err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}
	err = models.DeleteCategory(Id)
	if err != nil {
		log.Println("failed to delete category", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, nil)
}
