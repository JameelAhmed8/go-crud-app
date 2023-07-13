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

type ArticleUpdateReq struct {
	ArticleID  uuid.UUID `json:"article_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	CategoryID uuid.UUID `json:"category_id"`
}

func ArticleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	defer r.Body.Close()
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to parse request body: %v", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	body := &ArticleUpdateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		log.Printf("failed to parse JSON object: %v", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	err = models.UpdateArticle(body.ArticleID, body.Title, body.Content, body.Author, body.CategoryID)
	if err != nil {
		log.Println("failed to update article", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, nil)
}
