package repository

import (
	"fmt"
	"yokoso_api/model"

	"gorm.io/gorm"
)

// GetWords はデータベースからワード情報を取得して返す
func GetWords(db *gorm.DB, id int) ([]model.Word, error) {
	var words []model.Word

	// ワード情報を取得
	word := db.Where("category_id = ?", id).Find(&words)
	if word.Error != nil {
		return nil, fmt.Errorf("ワード情報の取得に失敗しました: %v", word.Error)
	}

	return words, nil
}
