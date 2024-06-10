package controller

import (
	"fmt"
	"net/http"
	"time"

	"yokoso_api/db"
	"yokoso_api/repository"

	"github.com/labstack/echo/v4"
)

func ApiFestival(c echo.Context) error {
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)
	festivals, festivalImgs, festivalDescriptions, err := repository.GetFestivals(dbCon)
	if err != nil {
		// エラーハンドリング
		fmt.Println("祭り情報の取得に失敗しました:", err)
		return err
	}

	// フェスティバル情報に画像と説明を追加
	type FestivalResponse struct {
		FestivalName        string   `json:"festival_name"`
		FestivalImgs        []string `json:"festival_imgs"`
		FestivalDate        string   `json:"festival_date"`
		FestivalStartTime   string   `json:"festival_start_time"`
		FestivalEndTime     string   `json:"festival_end_time"`
		FestivalPrefecture  string   `json:"festival_prefecture"`
		FestivalAddress     string   `json:"festival_address"`
		FestivalLatitude    string   `json:"festival_latitude"`
		FestivalLongitude   string   `json:"festival_longitude"`
		FestivalTitle       []string `json:"festival_title"`
		FestivalDescription []string `json:"festival_description"`
	}

	var response []FestivalResponse
	for _, festival := range festivals {
		var imgs []string
		var titles []string
		var descriptions []string

		//FestivalIDごとの画像を配列に追加
		for _, img := range festivalImgs {
			if img.FestivalID == festival.FestivalID {
				imgs = append(imgs, img.FestivalImg)
			}
		}

		//FestivalIDごとの見出しと説明を配列に追加
		for _, desc := range festivalDescriptions {
			if desc.FestivalID == festival.FestivalID {
				titles = append(titles, desc.FestivalTitle)
				descriptions = append(descriptions, desc.FestivalDescription)
			}
		}

		response = append(response, FestivalResponse{
			FestivalName:        festival.FestivalName,
			FestivalImgs:        imgs,
			FestivalDate:        festival.FestivalDate.Format(time.RFC3339),
			FestivalStartTime:   festival.FestivalStartTime.Format(time.RFC3339),
			FestivalEndTime:     festival.FestivalEndTime.Format(time.RFC3339),
			FestivalPrefecture:  festival.FestivalAddress, 
			FestivalAddress:     festival.FestivalAddress,
			FestivalLatitude:    fmt.Sprintf("%f", festival.FestivalLatitude),
			FestivalLongitude:   fmt.Sprintf("%f", festival.FestivalLongitude),
			FestivalTitle:       titles,
			FestivalDescription: descriptions,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"item": response})
}
