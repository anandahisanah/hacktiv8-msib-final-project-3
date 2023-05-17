package service

import (
	"hacktiv8-msib-final-project-3/dto"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/repository/categoryrepository"
)

type CategoryService interface {
	CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr)
	GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr)
	UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo categoryrepository.CategoryRepository
}

func NewCategoryService(categoryRepo categoryrepository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (c *categoryService) CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr) {
	category := payload.ToEntity()

	createdCategory, err := c.categoryRepo.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateCategoryResponse{
		ID:        createdCategory.ID,
		Type:      createdCategory.Type,
		CreatedAt: createdCategory.CreatedAt,
	}

	return response, nil
}

func (c *categoryService) GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr) {
	categories, err := c.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllCategoriesResponse{}
	for _, category := range categories {
		response = append(response, dto.GetAllCategoriesResponse{
			ID:        category.ID,
			Type:      category.Type,
			UpdatedAt: category.UpdatedAt,
			CreatedAt: category.CreatedAt,
		})
	}

	return response, nil
}

func (c *categoryService) UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	oldCategory, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	newCategory := payload.ToEntity()

	updatedCategory, err2 := c.categoryRepo.UpdateCategory(oldCategory, newCategory)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateCategoryResponse{
		ID:        updatedCategory.ID,
		Type:      updatedCategory.Type,
		UpdatedAt: updatedCategory.UpdatedAt,
	}

	return response, nil
}

func (c *categoryService) DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	category, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	if err := c.categoryRepo.DeleteCategory(category); err != nil {
		return nil, err
	}

	response := &dto.DeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}

	return response, nil
}
