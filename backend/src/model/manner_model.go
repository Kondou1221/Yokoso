package model

type Manner struct {
    MannerID          int    `gorm:"primaryKey;autoIncrement" json:"manner_id"`
    MannerTitle       string `gorm:"type:varchar(255);not null" json:"manner_title"`
    MannerDescription string `gorm:"type:varchar(500);not null" json:"manner_description"`
    CategoryID        int    `json:"category_id"`
}