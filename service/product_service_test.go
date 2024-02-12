package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productService = ProductService{Repository: productRepository}

	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productService = ProductService{Repository: productRepository}

	product := entity.Product{
		Id:   "2",
		Name: "Kaca mata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Kaca mata'")
	assert.Equal(t, &product, result, "result has to be product data with id '2'")
}

func TestProductServiceGetAllProductsNotFound(t *testing.T) {
	var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productService = ProductService{Repository: productRepository}

	productRepository.Mock.On("FindAllProducts").Return(nil)

	products, err := productService.GetAllProducts()

	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "products not found", err.Error(), "error response has to be 'products not found'")
}

func TestProductServiceGetAllProducts(t *testing.T) {
	var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productService = ProductService{Repository: productRepository}

	product1 := entity.Product{
		Id:   "1",
		Name: "BritAma",
	}

	product2 := entity.Product{
		Id:   "2",
		Name: "Simpedes",
	}

	products := []*entity.Product{&product1, &product2}

	productRepository.Mock.On("FindAllProducts").Return(products)

	result, err := productService.GetAllProducts()

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, products[0].Id, result[0].Id, "result has to be '1'")
	assert.Equal(t, products[0].Name, result[0].Name, "result has to be 'BritAma'")
	assert.Equal(t, products[1].Id, result[1].Id, "result has to be '2'")
	assert.Equal(t, products[1].Name, result[1].Name, "result has to be 'Simpedes'")
	assert.Equal(t, products[0], result[0], "result has to be product data with id '1'")
	assert.Equal(t, products[1], result[1], "result has to be product data with id '2'")
}

func TestProductServiceGetAllProductsNotFoundNotNil(t *testing.T) {
	var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productService = ProductService{Repository: productRepository}

	product1 := entity.Product{
		Id:   "1",
		Name: "BritAma",
	}

	product2 := entity.Product{
		Id:   "2",
		Name: "Simpedes",
	}

	products := []*entity.Product{&product1, &product2}

	productRepository.Mock.On("FindAllProducts").Return(products)

	products, err := productService.GetAllProducts()

	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "products not found", err.Error(), "error response has to be 'products not found'")
}
