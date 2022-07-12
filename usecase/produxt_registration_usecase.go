package usecase

import (
	"go-mongod/model"
	"go-mongod/repository"
)

type ProductRegistrationUseCase interface {
	Register(newProduct *model.Product) error
}

type productRegistrationUseCase struct {
	repo repository.ProductRepository
}

func (p *productRegistrationUseCase) Register(newProduct *model.Product) error {
	return p.repo.Add(newProduct)
}

func NewProductRegistrationUseCase(repo repository.ProductRepository) ProductRegistrationUseCase {
	return &productRegistrationUseCase{
		repo: repo,
	}
}