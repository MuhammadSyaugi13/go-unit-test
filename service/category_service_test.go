package service

import (
	"testing"
	"unit-test-go/entity"
	"unit-test-go/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	resultCategory, err := categoryService.Get("1")
	assert.Nil(t, resultCategory)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {

	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	// program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	resultCategory, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, resultCategory)
	assert.Equal(t, category.Id, resultCategory.Id)
	assert.Equal(t, category.Name, resultCategory.Name)
}
