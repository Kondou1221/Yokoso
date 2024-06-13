package model

type Word struct {
	WordID          int       `json:"word_id" gorm:"primaryKey;autoIncrement"`
	Scene           string    `json:"scene" gorm:"type:varchar(500);not null"`
	WordTranslation string    `json:"word_translation" gorm:"type:varchar(255);not null"`
	Pronunciation   string    `json:"pronunciation" gorm:"type:varchar(255);not null"`
	CategoryID      int       `json:"category_id"`
}
