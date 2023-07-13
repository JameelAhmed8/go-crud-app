package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Category struct {
	CategoryID uuid.UUID
	Name       string
}

const (
	dataFilePath  = "./pkg/db/txt/categories.txt"
	dataDelimiter = "|"
)

// CreateCategory creates a new category and saves it to the file
func CreateCategory(name string) (uuid.UUID, error) {
	categoryID := uuid.New()

	category := Category{
		CategoryID: categoryID,
		Name:       name,
	}

	file, err := os.OpenFile(dataFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	categoryString := fmt.Sprintf("%s%s%s\n", category.CategoryID.String(), dataDelimiter, category.Name)
	if _, err := file.WriteString(categoryString); err != nil {
		return uuid.Nil, fmt.Errorf("failed to write to data file: %v", err)
	}

	return categoryID, nil
}

// UpdateCategory updates the category with the specified ID and saves it to the file
func UpdateCategory(categoryID uuid.UUID, name string) error {
	file, err := os.OpenFile(dataFilePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, dataDelimiter)
		if len(fields) != 2 {
			continue
		}

		currID := uuid.MustParse(fields[0])
		if currID == categoryID {
			line = fmt.Sprintf("%s%s%s\n", categoryID.String(), dataDelimiter, name)
		}

		lines = append(lines, line)
	}

	if scanner.Err() != nil {
		return fmt.Errorf("error scanning data file: %v", scanner.Err())
	}

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate data file: %v", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek data file: %v", err)
	}

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			writer.Flush()
			return fmt.Errorf("failed to write to data file: %v", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush data file writer: %v", err)
	}

	return nil
}

// DeleteCategory deletes the category with the specified ID from the file
func DeleteCategory(categoryID uuid.UUID) error {
	file, err := os.OpenFile(dataFilePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, dataDelimiter)
		if len(fields) != 2 {
			continue
		}

		currID := uuid.MustParse(fields[0])
		if currID != categoryID {
			lines = append(lines, line)
		}
	}

	if scanner.Err() != nil {
		return fmt.Errorf("error scanning data file: %v", scanner.Err())
	}

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate data file: %v", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek data file: %v", err)
	}

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			writer.Flush()
			return fmt.Errorf("failed to write to data file: %v", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush data file writer: %v", err)
	}

	return nil
}

// ListCategories returns a list of all categories read from the file
func ListCategories() ([]Category, error) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var categories []Category

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, dataDelimiter)
		if len(fields) != 2 {
			continue
		}

		categoryID := uuid.MustParse(fields[0])
		name := fields[1]

		category := Category{
			CategoryID: categoryID,
			Name:       name,
		}

		categories = append(categories, category)
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("error scanning data file: %v", scanner.Err())
	}

	return categories, nil
}

// GetCategoryByID returns the category with the specified ID from the file
func GetCategoryByID(categoryID uuid.UUID) (*Category, error) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, dataDelimiter)
		if len(fields) != 2 {
			continue
		}

		currID := uuid.MustParse(fields[0])
		if currID == categoryID {
			category := Category{
				CategoryID: currID,
				Name:       fields[1],
			}
			return &category, nil
		}
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("error scanning data file: %v", scanner.Err())
	}

	return nil, fmt.Errorf("category not found")
}
