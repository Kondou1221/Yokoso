package repository

import (
	"fmt"
	"yokoso_api/model"

	"gorm.io/gorm"
)

// GetManners はデータベースからマナー情報を取得して返す
func GetManners(db *gorm.DB, id int) ([]model.Manner, error) {
	var manners []model.Manner

	// フェスティバル情報を取得
	manner := db.Where("category_id = ?", id).Find(&manners)
	if manner.Error != nil {
		return nil, fmt.Errorf("マナー情報の取得に失敗しました: %v", manner.Error)
	}

	return manners, nil
}

