package categoryrepository

import (
	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*entity.Category, errs.MessageErr)
	GetAllCategories() ([]entity.Category, errs.MessageErr)
	GetCategoryByID(id uint) (*entity.Category, errs.MessageErr)
	UpdateCategory(oldCategory *entity.Category, newCategory *entity.Category) (*entity.Category, errs.MessageErr)
}
