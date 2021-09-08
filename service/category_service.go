package service

import (
	"Bab_3_Unit-Testing/entity"
	"Bab_3_Unit-Testing/repository"
	"errors"
)

type CategoryService struct{
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	}else {
		return category, nil
	}
}