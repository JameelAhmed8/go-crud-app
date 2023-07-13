package models

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/JameelAhmed8/go-crud-app/pkg/config"
	"github.com/google/uuid"
)

var db *sql.DB

type Author struct {
	AuthorID    uuid.UUID
	Name        string
	Email       string
	Affiliation string
	Biography   string
	Website     string
}

func (tup *Author) allFields() []any {
	return []any{
		&tup.AuthorID,
		&tup.Name,
		&tup.Email,
		&tup.Affiliation,
		&tup.Biography,
		&tup.Website,
	}
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func CreateAuthor(tup *Author) (uuid.UUID, error) {
	var authorId uuid.UUID
	res := db.QueryRow(
		`INSERT INTO Author(
			author_id, 
			name, 
			email, 
			affiliation, 
			biography, 
			website
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING author_id`,
		tup.AuthorID,
		tup.Name,
		tup.Email,
		tup.Affiliation,
		tup.Biography,
		tup.Website,
	)
	if err := res.Scan(&authorId); err != nil {
		log.Printf("failed to scan row: %v", err)
		return authorId, err
	}
	return authorId, nil
}

func ListAuthors() ([]*Author, error) {
	rows, err := db.Query(
		`SELECT
			author_id, 
			name, 
			email, 
			affiliation, 
			biography, 
			website
		FROM author`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tups := make([]*Author, 0)
	for rows.Next() {
		tup := Author{}
		if err := rows.Scan(tup.allFields()...); err != nil {
			log.Printf("failed to scan row: %v", err)
			continue
		}
		tups = append(tups, &tup)
	}
	return tups, rows.Err()
}

func GetAuthor(id uuid.UUID) (*Author, error) {
	tup := Author{}
	err := db.QueryRow(
		`SELECT
			author_id, 
			name, 
			email, 
			affiliation, 
			biography, 
			website
		FROM author
		WHERE author_id = $1`,
		id,
	).Scan(
		&tup.AuthorID,
		&tup.Name,
		&tup.Email,
		&tup.Affiliation,
		&tup.Biography,
		&tup.Website)
	if err != nil {

		return nil, err
	}

	return &tup, nil
}

type AuthorUpdateTup struct {
	AuthorId uuid.UUID
	Fields   []byte
}

func UpdateAuthor(authorID uuid.UUID, fields map[string]interface{}) error {
	// Construct the UPDATE query dynamically
	query := "UPDATE author SET"
	values := make([]interface{}, 0)

	i := 1 // Parameter index
	for field, value := range fields {
		query += " " + field + " = $" + strconv.Itoa(i) + ","
		values = append(values, value)
		i++
	}

	// Remove the trailing comma and add the WHERE clause
	query = strings.TrimSuffix(query, ",") + " WHERE author_id = $" + strconv.Itoa(i)

	// Append the authorID as the last value
	values = append(values, authorID)

	// Execute the update query
	_, err := db.Exec(query, values...)
	if err != nil {
		log.Printf("failed to update author: %v", err)
		return err
	}

	return nil
}

func DeleteAuthor(authorID uuid.UUID) error {
	query := "DELETE FROM Author WHERE author_id = $1"
	_, err := db.Exec(query, authorID)
	if err != nil {
		log.Printf("failed to delete author: %v", err)
		return err
	}

	return nil
}
