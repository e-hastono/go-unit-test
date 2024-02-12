package repository

import (
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := repository.Mock.Called(id)

	argument0 := arguments.Get(0)

	if argument0 == nil {
		return nil
	}

	product := argument0.(entity.Product)

	return &product
}

func (repository *ProductRepositoryMock) FindAllProducts() []*entity.Product {
	arguments := repository.Mock.Called()

	argument0 := arguments.Get(0)

	if argument0 == nil {
		return nil
	}

	products := argument0.([]*entity.Product)

	return products
}
