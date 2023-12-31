package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

/* ====================================== */

type ICategoryController interface {
	CreateCategory(c *gin.Context)
	GetCategoryInfo(c *gin.Context)
	GetPrimaryCategories(c *gin.Context)
	GetSecondaryCategories(c *gin.Context)
	GetCategoryList(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController() *CategoryController {
	categoryService := service.NewCategoryService()
	return &CategoryController{categoryService}
}

/* ====================================== */

// 添加一级分类
func (cc *CategoryController) CreatePrimaryCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	code := errmsg.SUCCESS
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		category.ParentID = nil
		code = cc.categoryService.CreateCategory(&category)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    &category,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 添加二级分类
func (cc *CategoryController) CreateSecondaryCategory(c *gin.Context) {
	parentID, _ := strconv.Atoi(c.Param("id"))
	var category model.Category
	err := c.ShouldBindJSON(&category)
	code := errmsg.SUCCESS
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		category.ParentID = &parentID
		code = cc.categoryService.CreateCategory(&category)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    &category,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询分类信息
func (cc *CategoryController) GetCategoryInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := cc.categoryService.GetCategoryInfo(uint(id))

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func (cc *CategoryController) GetPrimaryCategories(c *gin.Context) {
	data, total, code := cc.categoryService.GetPrimaryCategories()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (cc *CategoryController) GetSecondaryCategories(c *gin.Context) {
	parentID, _ := strconv.Atoi(c.Param("id"))
	data, total, code := cc.categoryService.GetSecondaryCategories(parentID)
	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"parent_cid": parentID,
		"data":       data,
		"total":      total,
		"message":    errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func (cc *CategoryController) GetCategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := cc.categoryService.GetCategoryList(pageSize, pageNum)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 编辑分类名
func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var category model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&category)

	code := cc.categoryService.UpdateCategory(uint(id), &category)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    &category,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 删除分类名
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := cc.categoryService.DeleteCategory(uint(id))

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
