package controller

import (
	"fmt"
	"net/http"

	"yokoso_api/db"
	"yokoso_api/repository"

	"github.com/labstack/echo/v4"
)

// CategoryResponse defines the structure for the JSON response
type CategoryResponse struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func ApiCategory(c echo.Context) error {
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)
	categories, err := repository.GetCategories(dbCon)
	if err != nil {
		// エラーハンドリング
		fmt.Println("カテゴリの取得に失敗しました:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "カテゴリの取得に失敗しました"})
	}

	// Create a slice of CategoryResponse for the API response
	categoriesAPI := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		categoriesAPI[i] = CategoryResponse{
			CategoryID:   category.CategoryID,
			CategoryName: category.CategoryName,
		}
	}

	return c.JSON(http.StatusOK, categoriesAPI)
}
