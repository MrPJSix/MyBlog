package repository

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type ICategoryRepo interface {
	CheckByName(name string) int
	Create(category *model.Category) int
	GetInfo(id uint) (model.Category, int)
	GetList(pageSize, offset int) ([]model.Category, int64, int)
	Update(id uint, category *model.Category) int
	Delete(id uint) int
}

type CategoryRepo struct{}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{}
}

/* ====================================== */

// 检查标签名是否存在
func (cr *CategoryRepo) CheckByName(name string) int {
	var cate model.Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 新增分类
func (cr *CategoryRepo) Create(category *model.Category) int {
	code := cr.CheckByName(category.Name)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Create(category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个分类信息
func (cr *CategoryRepo) GetInfo(id uint) (model.Category, int) {
	var cate model.Category
	err := db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return cate, errmsg.ERROR_CATE_NOT_EXIST
		}
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCESS
}

// 查询分类列表
func (cr *CategoryRepo) GetList(pageSize, offset int) ([]model.Category, int64, int) {
	var cateList []model.Category
	var total int64

	err := db.Limit(pageSize).Offset(offset).Find(&cateList).Error
	db.Model(&cateList).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errmsg.ERROR
	}
	return cateList, total, errmsg.SUCCESS
}

func (cr *CategoryRepo) Update(id uint, category *model.Category) int {
	if code := cr.CheckByName(category.Name); code != errmsg.SUCCESS {
		return code
	}
	var cate model.Category
	var maps = make(map[string]interface{})
	maps["name"] = category.Name
	log.Println(maps)
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (cr *CategoryRepo) Delete(id uint) int {
	var cate model.Category
	_, code := cr.GetInfo(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
