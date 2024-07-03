package model

import (
	"time"
)
type WeekName struct {
	WeekNameID    int    `gorm:"primaryKey;autoIncrement" json:"week_name_id"`
	DayID      	  int    `json:"day_id"`
    PlaceID       int    `json:"place_id"`
    StartTime 	  time.Time   `gorm:"not null;type:time" json:"start_time"`
	FinishTime 	  time.Time   `gorm:"not null;type:time" json:"finish_time"`
}