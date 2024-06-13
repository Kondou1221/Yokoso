package model

type FestivalImg struct {
	FestivalImgId    	int      `gorm:"primaryKey;autoIncrement" json:"festival_img_id"`
    FestivalImg      	string   `gorm:"type:varchar(255)" json:"festival_img"`
	FestivalID  		int   	 `json:"festival_id"`
}