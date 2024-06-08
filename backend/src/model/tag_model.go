package model

type Tag struct {
	TagID    int    `gorm:"primaryKey;autoIncrement" json:"tag_id"`
    TagName  int 	`gorm:"type:varchar(255)" json:"tag_name"`
	TagMap   TagMap `gorm:"foreignKey:TagID"`
}