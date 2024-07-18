package repository

import (
	"fmt"
	"yokoso_api/model"

	"gorm.io/gorm"
)

// GetFestivals はデータベースから祭り情報を取得して返す
func GetFestivals(db *gorm.DB) ([]model.Festival, []model.FestivalImg, []model.FestivalDescription, error) {
	var festivals []model.Festival
	var festivalImgs []model.FestivalImg
	var festivalDescriptions []model.FestivalDescription

	// フェスティバル情報を取得
	fes := db.Select("festival_id", "festival_name", "festival_prefecture", "festival_date", "festival_address", "festival_start_time", "festival_end_time", "festival_latitude", "festival_longitude").Order("festival_id asc").Find(&festivals)
	if fes.Error != nil {
		return nil, nil, nil, fmt.Errorf("祭り情報の取得に失敗しました: %v", fes.Error)
	}

	// フェスティバル画像情報を取得
	img := db.Select("festival_img_id", "festival_img", "festival_id").Order("festival_img_id asc").Find(&festivalImgs)
	if img.Error != nil {
		return nil, nil, nil, fmt.Errorf("祭り画像情報の取得に失敗しました: %v", img.Error)
	}

	// フェスティバル説明情報を取得
	desc := db.Select("festival_title_id", "festival_title", "festival_description", "festival_id").Order("festival_title_id asc").Find(&festivalDescriptions)
	if desc.Error != nil {
		return nil, nil, nil, fmt.Errorf("祭り説明情報の取得に失敗しました: %v", desc.Error)
	}

	return festivals, festivalImgs, festivalDescriptions, nil
}
