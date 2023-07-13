package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type CategoryCreateReq struct {
	Name string `json:"category_name"`
}

type CategoryCreateResp struct {
	ID uuid.UUID `json:"category_id"`
}

func CategoryCreateHandler(w http.ResponseWriter, r *http.Request) {

	// Parse request body.
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		log.Println("failed to read request body", err)
		return
	}
	defer r.Body.Close()

	body := &CategoryCreateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, nil)
		log.Println("failed to parse JSON object", err)
		return
	}

	catId, err := models.CreateCategory(body.Name)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		log.Println("failed to add/create author", err)
		return
	}

	// Create response.
	resp := &CategoryCreateResp{
		ID: catId,
	}

	utils.WriteJSONResponse(w, http.StatusOK, resp)
}
