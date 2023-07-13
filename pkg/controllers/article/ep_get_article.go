package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type ArticleGetResp struct {
	ArticleID   uuid.UUID `json:"article_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishDate string    `json:"publish_date"`
	Author      string    `json:"author"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func ArticleGetHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	articleIDStr := qp.Get("article_id")
	articleID, err := uuid.Parse(articleIDStr)
	if err != nil {
		log.Printf("failed to parse articleIDStr=%v: %v", articleIDStr, err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}
	article, err := models.GetArticleByID(articleID)
	if err != nil {
		log.Println("failed to get article", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// Convert to endpoint output format
	data := ArticleGetResp{
		ArticleID:   article.ArticleID,
		Title:       article.Title,
		Content:     article.Content,
		PublishDate: article.PublishDate.String(),
		Author:      article.Author,
		CategoryID:  article.CategoryID,
	}

	utils.WriteJSONResponse(w, http.StatusOK, data)
}
