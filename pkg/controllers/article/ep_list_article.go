package controllers

import (
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

type ArticlesListResp struct {
	Articles []*ArticleListRespObj `json:"articles"`
}

type ArticleListRespObj struct {
	ArticleID   uuid.UUID `json:"article_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishDate string    `json:"publish_date"`
	Author      string    `json:"author"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func ArticlesListHandler(w http.ResponseWriter, r *http.Request) {

	articles, err := models.ListArticles()
	if err != nil {
		log.Println("failed to list articles", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// Convert to endpoint output format
	data := make([]*ArticleListRespObj, 0, len(articles))
	for _, article := range articles {
		data = append(data, &ArticleListRespObj{
			ArticleID:   article.ArticleID,
			Title:       article.Title,
			Content:     article.Content,
			PublishDate: article.PublishDate.String(),
			Author:      article.Author,
			CategoryID:  article.CategoryID,
		})
	}
	utils.WriteJSONResponse(w, http.StatusOK, data)
}
