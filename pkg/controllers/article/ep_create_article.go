package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type ArticleCreateReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	CategoryID string `json:"category_id"`
}

type ArticleCreateResp struct {
	ID string `json:"article_id"`
}

func ArticleCreateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body.
	body := &ArticleCreateReq{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, nil)
		log.Println("failed to parse JSON object", err)
		return
	}
	defer r.Body.Close()

	categoryID, err := uuid.Parse(body.CategoryID)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, nil)
		log.Println("failed to parse category ID", err)
		return
	}

	articleID, err := models.CreateArticle(body.Title, body.Content, body.Author, categoryID)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		log.Println("failed to create article", err)
		return
	}

	// Create response.
	resp := &ArticleCreateResp{
		ID: articleID.String(),
	}

	utils.WriteJSONResponse(w, http.StatusOK, resp)
}
