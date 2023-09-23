package model

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(50);not null" json:"name"`
}
