package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ArticleID   uuid.UUID
	Title       string
	Content     string
	PublishDate time.Time
	Author      string
	CategoryID  uuid.UUID
}

var (
	articles []Article
)

// CreateArticle creates a new article and stores it in memory
func CreateArticle(title, content, author string, categoryID uuid.UUID) (uuid.UUID, error) {
	articleID := uuid.New()

	article := Article{
		ArticleID:   articleID,
		Title:       title,
		Content:     content,
		PublishDate: time.Now(),
		Author:      author,
		CategoryID:  categoryID,
	}

	articles = append(articles, article)

	return articleID, nil
}

// UpdateArticle updates the article with the specified ID in memory
func UpdateArticle(articleID uuid.UUID, title, content, author string, categoryID uuid.UUID) error {
	for i := range articles {
		if articles[i].ArticleID == articleID {
			articles[i].Title = title
			articles[i].Content = content
			articles[i].Author = author
			articles[i].CategoryID = categoryID
			return nil
		}
	}

	return fmt.Errorf("article not found")
}

// DeleteArticle deletes the article with the specified ID from memory
func DeleteArticle(articleID uuid.UUID) error {
	for i, article := range articles {
		if article.ArticleID == articleID {
			articles = append(articles[:i], articles[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("article not found")
}

// ListArticles returns a list of all articles stored in memory
func ListArticles() ([]Article, error) {
	return articles, nil
}

// GetArticleByID returns the article with the specified ID from memory
func GetArticleByID(articleID uuid.UUID) (*Article, error) {
	for _, article := range articles {
		if article.ArticleID == articleID {
			return &article, nil
		}
	}

	return nil, fmt.Errorf("article not found")
}

// GetArticlesByCategory returns a list of articles filtered by category from memory
func GetArticlesByCategory(categoryID uuid.UUID) ([]Article, error) {
	var filteredArticles []Article

	for _, article := range articles {
		if article.CategoryID == categoryID {
			filteredArticles = append(filteredArticles, article)
		}
	}

	return filteredArticles, nil
}
