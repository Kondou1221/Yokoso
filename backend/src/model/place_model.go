package model

type Place struct {
    PlaceID           int         `gorm:"primaryKey;autoIncrement" json:"place_id"`
    PlaceName         string      `gorm:"type:varchar(255)" json:"place_name"`
    PlaceDescription  string      `gorm:"type:varchar(500)" json:"place_description"`
    PlaceLatitude     float64     `gorm:"type:double; not null; index" json:"place_latitude"`
    PlaceLongitude    float64     `gorm:"type:double; not null; index" json:"place_longitude"`
    LowestPrice       int         `json:"lowest_price"`
    HighestPrice      int         `json:"highest_price"`
    SubCategoryID     int         `json:"sub_category_id"`
    PlaceImage        PlaceImage  `gorm:"foreignKey:PlaceID"`
    WeekName          WeekName    `gorm:"foreignKey:PlaceID"`
    TagMap            TagMap      `gorm:"foreignKey:PlaceID"`
}