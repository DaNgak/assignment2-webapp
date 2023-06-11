package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryRepo service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add category success"})
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var updateCategory model.Category

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid request data"})
		return
	}

	// Convert string to int
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Category ID"})
		return
	}

	err = ct.categoryService.Update(id, updateCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "category update success"})
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Convert string to int
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid Category ID"})
		return
	}

	err = ct.categoryService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "category delete success"})
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid category ID"})
		return
	}

	category, err := ct.categoryService.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	categoryList, err := ct.categoryService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoryList)
}