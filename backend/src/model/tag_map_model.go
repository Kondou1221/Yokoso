package model

type TagMap struct {
	TagID    int    `gorm:"primaryKey;" json:"tag_id"`
    PlaceID  int    `json:"place_id"` 	
}