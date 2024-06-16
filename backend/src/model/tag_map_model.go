package model

type TagMap struct {
	TagMap    int    `gorm:"primaryKey;autoIncrement" json:"tag_map"`
	TagID    int    `json:"tag_id"`
    PlaceID  int    `json:"place_id"` 	
}