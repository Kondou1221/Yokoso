package repository

import (
	"fmt"
	"yokoso_api/model"

	"gorm.io/gorm"
)

// GetCategories はデータベースからカテゴリ情報を取得して返す
func GetCategories(db *gorm.DB) ([]model.Category, error) {
	var categories []model.Category
	// 必要なフィールドのみを選択
	result := db.Select("category_id", "category_name").Order("category_id asc").Find(&categories)
	if result.Error != nil {
		return nil, fmt.Errorf("カテゴリ情報の取得に失敗しました: %v", result.Error)
	}

	return categories, nil
}
