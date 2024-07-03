package repository

import (
	"fmt"
	"yokoso_api/model"

	"gorm.io/gorm"
)

// GetMaps はデータベースからマップ情報を取得して返す
func GetMaps(db *gorm.DB) ([]model.Place, []model.PlaceImage, []model.Tag, []model.WeekName, []model.Subcategory, []model.TagMap, error) {
	var places []model.Place
	var placeImages []model.PlaceImage
	var tags []model.Tag
	var weekNames []model.WeekName
	var subCategories []model.Subcategory
	var tagMaps []model.TagMap

	// マップ情報を取得
	if err := db.Select("place_id", "place_name", "place_description", "place_latitude", "place_longitude", "lowest_price", "highest_price", "sub_category_id").Order("place_id asc").Find(&places).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("マップ情報の取得に失敗しました: %v", err)
	}

	// マップ画像情報を取得
	if err := db.Select("place_img_id", "place_id", "place_img").Order("place_img_id asc").Find(&placeImages).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("マップ画像情報の取得に失敗しました: %v", err)
	}

	// タグ情報を取得
	if err := db.Select("tag_id", "tag_name").Order("tag_id asc").Find(&tags).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("タグ情報の取得に失敗しました: %v", err)
	}

	// 曜日情報を取得
	if err := db.Select("day_id", "place_id", "start_time", "finish_time").Order("week_name_id asc").Find(&weekNames).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("曜日情報の取得に失敗しました: %v", err)
	}

	// サブカテゴリ情報を取得
	if err := db.Select("sub_category_id", "sub_category_name", "category_id").Order("sub_category_id asc").Find(&subCategories).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("サブカテゴリ情報の取得に失敗しました: %v", err)
	}

	// タグマップ情報を取得
	if err := db.Select("tag_id", "place_id").Order("tag_id asc").Find(&tagMaps).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("タグマップ情報の取得に失敗しました: %v", err)
	}

	return places, placeImages, tags, weekNames, subCategories, tagMaps, nil
}
