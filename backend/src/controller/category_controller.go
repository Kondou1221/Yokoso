package controller

import (
	"fmt"
	"net/http"

	"yokoso_api/db"
	"yokoso_api/model"
	"yokoso_api/repository"

	"github.com/labstack/echo/v4"
)

func ApiCategory(c echo.Context) error {
	db := database.Connect()
	defer db.Close()
	categories, err := repository.GetCategories(db)
	
	if err != nil {
		// エラーハンドリング
		fmt.Println("カテゴリの取得に失敗しました:", err)
		return err
	}

	categoriesAPI := make([]model.Category, 0)
	for _, category := range categories {
		categoryAPI := model.Category{
			CategoryID:   category.CategoryID,
			CategoryName: category.CategoryName,
		}
		categoriesAPI = append(categoriesAPI, categoryAPI)
	}

	return c.JSON(http.StatusOK, categoriesAPI)
}