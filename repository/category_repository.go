package repository

import "Bab_3_Unit-Testing/entity"

type CategoryRepository interface{
	FindById(id string) *entity.Category
}