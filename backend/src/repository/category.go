// repository/category.go
package repository

import (
	"yokoso_api/model"
	"database/sql"
	"fmt"
)

// GetCategories はデータベースからカテゴリ情報を取得して返す
func GetCategories(db *sql.DB) ([]model.Category, error) {
	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		return nil, fmt.Errorf("カテゴリ情報の取得に失敗しました: %v", err)
	}
	defer rows.Close()

	var categories []model.Category

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			return nil, fmt.Errorf("カテゴリ情報のスキャンに失敗しました: %v", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("カテゴリ情報の取得中にエラーが発生しました: %v", err)
	}

	return categories, nil
}
