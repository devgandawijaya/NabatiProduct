package service

import (
	"go-api/model"
	"go-api/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllProduct() ([]model.ProductCategoryWithProducts, error) {
	return s.Repo.GetAllProduct()
}
