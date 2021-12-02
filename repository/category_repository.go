package repository

import "github.com/akbarismail/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
