package model

type PlaceImage struct {
	PlaceImgID    int    `gorm:"primaryKey;autoIncrement" json:"place_img_id"`
    PlaceID       int 	 `gorm:"index" json:"place_id"`
    PlaceImg 	  string `gorm:"type:varchar(255); not null" json:"place_img"`
 }