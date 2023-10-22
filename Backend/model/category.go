package model

type Category struct {
	ID             int        `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name           string     `gorm:"type:varchar(50);not null;comment:分类名" json:"name"`
	ParentID       *int       `gorm:"comment:父类分类ID" json:"parent_id"`
	ParentCategory *Category  `gorm:"foreignKey:ParentID" json:"parent_category"`
	SubCategories  []Category `gorm:"foreignKey:ParentID" json:"sub_categories"`
}
