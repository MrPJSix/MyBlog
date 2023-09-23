package repository

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

// 查看分类是否存在
func CheckCategoryByName(name string) int {
	var cate model.Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

func CheckCategoryByID(id int) int {
	err = db.Where("id = ?", id).Find(&model.Category{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_CATE_NOT_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCate(data *model.Category) int {
	code := CheckCategoryByName(data.Name)
	if code != errmsg.SUCCESS {
		return code
	}
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个分类信息
func GetCateInfo(id int) (model.Category, int) {
	var cate model.Category
	err = db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return cate, errmsg.ERROR_CATE_NOT_EXIST
		}
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCESS
}

// 查询分类列表
func GetCate(pageSize, pageNum int) ([]model.Category, int64) {
	var cateList []model.Category
	var total int64

	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = -1
	}

	offset := (pageNum - 1) * pageSize
	if pageNum == 0 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&cateList).Error
	db.Model(&cateList).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cateList, total
}

// 编辑分类信息
func EditCate(id int, data *model.Category) int {
	if code := CheckCategoryByName(data.Name); code != errmsg.SUCCESS {
		return code
	}
	var cate model.Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	log.Println(maps)
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCate(id int) int {
	var cate model.Category
	_, code := GetCateInfo(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
