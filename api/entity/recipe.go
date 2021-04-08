package entity

type Recipe struct {
	ID           int    `gorm:"primary_key;auto_increment" json:"id"`
	Image        string `gorm:"type:varchar(255);not null" json:"image"`
	RecipeName   string `gorm:"type:varchar(255);not null" json:"name"`
	Causions     string `gorm:"type:varchar(255);not null" json:"causions"`
	Instructions string `gorm:"type:varchar(255);not null" json:"instructions"`
	Calories     int    ` json:"calories"`
}
