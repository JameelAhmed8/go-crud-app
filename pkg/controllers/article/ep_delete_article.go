package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JameelAhmed8/go-crud-app/pkg/models"
	"github.com/JameelAhmed8/go-crud-app/pkg/utils"
	"github.com/google/uuid"
)

func ArticleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	articleIDStr := qp.Get("article_id")
	fmt.Println(articleIDStr)
	articleID, err := uuid.Parse(articleIDStr)
	if err != nil {
		log.Printf("failed to parse articleIDStr=%v: %v", articleIDStr, err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	err = models.DeleteArticle(articleID)
	if err != nil {
		log.Println("failed to delete article", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, nil)
}
