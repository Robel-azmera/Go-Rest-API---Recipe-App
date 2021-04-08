package entity

type Party struct {
	ID          int   `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Leader       string `gorm:"type:varchar(255);not null" json:"leader"`
	Region       string `gorm:"type:varchar(255);not null" json:"region"`
	Vote        int ` json:"vote"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	HPRMembers        int ` json:"hpr"`

}
