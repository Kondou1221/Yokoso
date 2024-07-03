package model

type Category struct {
	CategoryID   int    `json:"category_id" gorm:"primaryKey"`
	CategoryName string `json:"category_name" gorm:"not null;unique"`
	Word         Word    `gorm:"foreignKey:CategoryID"`
	Manner       Manner  `gorm:"foreignKey:CategoryID"`
	Subcategory      Subcategory `gorm:"foreignKey:CategoryID"`
}
