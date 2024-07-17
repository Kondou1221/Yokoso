package model

import(
	"time"
)

type Festival struct {
    FestivalID        			int         		`gorm:"primaryKey;autoIncrement" json:"festival_id"`
    FestivalName      			string      		`gorm:"type:varchar(255); not null" json:"festival_name"`
	FestivalPrefecture          string              `gorm:"type:varchar(255); not null" json:"festival_prefecture"`
    FestivalDate      			time.Time   		`gorm:"type:date; not null" json:"festival_date"`
	FestivalAddress   			string      		`gorm:"type:varchar(255); not null" json:"festival_address"`
	FestivalStartTime 			time.Time   		`gorm:"type:time; not null" json:"festival_start_time"`
	FestivalEndTime   			time.Time   		`gorm:"type:time; not null" json:"festival_end_time"`
	FestivalLatitude  			float64     		`gorm:"type:double; not null" json:"festival_latitude"`
	FestivalLongitude 			float64      		`gorm:"type:double; not null" json:"festival_longitude"`
	FestivalDescription 		FestivalDescription `gorm:"foreignKey:FestivalID"`
	FestivalImg          		FestivalImg			`gorm:"foreignKey:FestivalID"`
}