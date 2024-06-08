package model

type Subcategory struct {
    SubCategoryID   int    `gorm:"primaryKey;autoIncrement" json:"sub_category_id"`
    SubCategoryName string `gorm:"type:varchar(255);not null;unique" json:"sub_category_name"`
    CategoryID      int    `json:"category_id"`
    Place           Place  `gorm:"foreignKey:SubCategoryID"`
}