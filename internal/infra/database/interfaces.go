package database

import "github.com/misterclayt0n/go-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id int) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
