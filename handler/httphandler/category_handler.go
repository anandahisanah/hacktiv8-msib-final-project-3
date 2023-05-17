package httphandler

import (
	"hacktiv8-msib-final-project-3/dto"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (c *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var requestBody dto.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdCategory, err := c.categoryService.CreateCategory(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdCategory)
}

func (c *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	categoryID := ctx.Param("categoryID")
	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Category id should be in unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var requestBody dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedCategory, err2 := c.categoryService.UpdateCategory(uint(categoryIDUint), &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

func (c *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryID := ctx.Param("categoryID")
	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Category id should be in unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err2 := c.categoryService.DeleteCategory(uint(categoryIDUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
