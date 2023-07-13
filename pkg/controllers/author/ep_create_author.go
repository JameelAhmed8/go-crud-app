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

type AuthorCreateReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Affiliation string `json:"affiliation"`
	Biography   string `json:"biography"`
	Website     string `json:"website"`
}

type AuthorCreateResp struct {
	ID uuid.UUID `json:"auth_id"`
}

func AuthorCreateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body.
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		log.Println("failed to read request body", err)
		return
	}
	defer r.Body.Close()

	body := &AuthorCreateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, nil)
		log.Println("failed to parse JSON object", err)
		return
	}

	authorTup := &models.Author{
		AuthorID:    uuid.New(),
		Name:        body.Name,
		Email:       body.Email,
		Affiliation: body.Affiliation,
		Biography:   body.Biography,
		Website:     body.Website,
	}
	authorId, err := models.CreateAuthor(authorTup)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		log.Println("failed to add/create author", err)
		return
	}

	// Create response.
	resp := &AuthorCreateResp{
		ID: authorId,
	}

	utils.WriteJSONResponse(w, http.StatusOK, resp)
}
