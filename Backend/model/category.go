package model

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name string `gorm:"type:varchar(50);not null;comment:分类名" json:"name"`
}
