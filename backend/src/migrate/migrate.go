package main

import (
	"fmt"
	"yokoso_api/db"
	"yokoso_api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.Category{},
		&model.Word{}, 
		&model.Manner{}, 
		&model.Subcategory{}, 
		&model.Place{}, 
		&model.PlaceImage{}, 
		&model.WeekName{},
		&model.Tag{},
		&model.TagMap{},
		&model.Festival{},
		&model.FestivalDescription{},
		&model.FestivalImg{},
	)
}