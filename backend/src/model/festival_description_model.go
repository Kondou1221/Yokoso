package model

type FestivalDescription struct {
	FestivalTitleID    	int      `gorm:"primaryKey;autoIncrement" json:"festival_title_id"`
    FestivalTitle      	string   `gorm:"type:varchar(255)" json:"festival_title"`
    FestivalDescription string   `gorm:"type:varchar(500)" json:"festival_description"`
	FestivalID  		int   	 `json:"festival_id"`
}