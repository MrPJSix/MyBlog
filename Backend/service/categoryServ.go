package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
)

/* ====================================== */

type ICategoryService interface {
	CreateCategory(category *model.Category) int
	GetCategoryInfo(id uint) (model.Category, int)
	GetPrimaryCategories() ([]model.Category, int64, int)
	GetSecondaryCategories(parentID int) ([]model.Category, int64, int)
	GetCategoryList(pageSize, pageNum int) ([]model.Category, int64, int)
	UpdateCategory(id uint, category *model.Category) int
	DeleteCategory(id uint) int
}

type CategoryService struct {
	categoryRepo *repository.CategoryRepo
}

func NewCategoryService() *CategoryService {
	categoryRepo := repository.NewCategoryRepo()
	return &CategoryService{categoryRepo}
}

/* ====================================== */

func (cs *CategoryService) CreateCategory(category *model.Category) int {
	return cs.categoryRepo.Create(category)
}

func (cs *CategoryService) GetCategoryInfo(id uint) (model.Category, int) {
	return cs.categoryRepo.GetInfo(id)
}

func (cs *CategoryService) GetPrimaryCategories() ([]model.Category, int64, int) {
	return cs.categoryRepo.GetPrimary()
}
func (cs *CategoryService) GetSecondaryCategories(parentID int) ([]model.Category, int64, int) {
	return cs.categoryRepo.GetSecondary(parentID)
}

func (cs *CategoryService) GetCategoryList(pageSize, pageNum int) ([]model.Category, int64, int) {
	var offset int
	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = -1
	}
	if pageNum <= 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	return cs.categoryRepo.GetList(pageSize, offset)
}

func (cs *CategoryService) UpdateCategory(id uint, category *model.Category) int {
	return cs.categoryRepo.Update(id, category)
}

func (cs *CategoryService) DeleteCategory(id uint) int {
	return cs.categoryRepo.Delete(id)
}
