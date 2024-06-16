package controller

import (
	"fmt"
	"net/http"
	"yokoso_api/db"
	"yokoso_api/repository"

	"github.com/labstack/echo/v4"
)

func ApiMap(c echo.Context) error {
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)

	places, placeImgs, tags, weekNames, subCategories, tagMaps, err := repository.GetMaps(dbCon)
	if err != nil {
		// エラーハンドリング
		fmt.Println("マップ情報の取得に失敗しました:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "マップ情報の取得に失敗しました"})
	}

	type Week struct {
		WeekNameID int    `json:"week_name_id"`
		StartTime  string `json:"start_time"`
		FinishTime string `json:"finish_time"`
	}

	type PlaceResponse struct {
		PlaceID          int      `json:"place_id"`
		PlaceName        string   `json:"place_name"`
		PlaceDescription string   `json:"place_description"`
		PlaceImgs        []string `json:"place_img"`
		Tags             []string `json:"tag"`
		Weeks            []Week   `json:"week"`
		PlaceLatitude    float64  `json:"place_latitude"`
		PlaceLongitude   float64  `json:"place_longitude"`
		LowestPrice      int      `json:"lowest_price"`
		HighestPrice     int      `json:"highest_price"`
		SubCategoryName  string   `json:"sub_category_name"`
	}

	var response []PlaceResponse

	for _, place := range places {
		var imgs []string
		for _, img := range placeImgs {
			if img.PlaceID == place.PlaceID {
				imgs = append(imgs, img.PlaceImg)
			}
		}

		var placeTags []string
		for _, tagMap := range tagMaps {
			if tagMap.PlaceID == place.PlaceID {
				for _, tag := range tags {
					if tag.TagID == tagMap.TagID {
						placeTags = append(placeTags, tag.TagName)
					}
				}
			}
		}

		var weeks []Week
		for _, week := range weekNames {
			if week.PlaceID == place.PlaceID {
				weeks = append(weeks, Week{
					WeekNameID: week.DayID,
					StartTime:  week.StartTime.Format("15:04"),
					FinishTime: week.FinishTime.Format("15:04"),
				})
			}
		}

		var subCategoryName string
		for _, subCategory := range subCategories {
			if subCategory.SubCategoryID == place.SubCategoryID {
				subCategoryName = subCategory.SubCategoryName
				break
			}
		}

		response = append(response, PlaceResponse{
			PlaceID:          place.PlaceID,
			PlaceName:        place.PlaceName,
			PlaceDescription: place.PlaceDescription,
			PlaceImgs:        imgs,
			Tags:             placeTags,
			Weeks:            weeks,
			PlaceLatitude:    place.PlaceLatitude,
			PlaceLongitude:   place.PlaceLongitude,
			LowestPrice:      place.LowestPrice,
			HighestPrice:     place.HighestPrice,
			SubCategoryName:  subCategoryName,
		})
	}

	return c.JSON(http.StatusOK, response)
}
