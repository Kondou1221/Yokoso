package model

type Subcategory struct {
    SubCategoryID   int    `json:"sub_category_id" gorm:"primaryKey;autoIncrement"`
    SubCategoryName string `json:"sub_category_name" gorm:"type:varchar(255);not null;unique"`
    CategoryID      int    `json:"category_id"`
    Place           Place  `gorm:"foreignKey:SubCategoryID"`
}