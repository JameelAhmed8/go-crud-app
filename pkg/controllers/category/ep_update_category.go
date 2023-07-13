package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type CategoryUpdateReq struct {
	Id   uuid.UUID `json:"category_id"`
	Name string    `json:"name"`
}

func CategoryUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse req body.
	defer r.Body.Close()
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to parse req body: %v", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	body := &CategoryUpdateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		log.Printf("failed to parse JSON object: %v", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	err = models.UpdateCategory(body.Id, body.Name)
	if err != nil {
		log.Println("failed to update category", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, nil)
}
