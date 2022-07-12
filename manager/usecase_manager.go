package manager

import "go-mongod/usecase"

type UseCaseManager interface {
	ProductRepo() usecase.ProductRegistrationUseCase
}

type usecaseManager struct {
	repositoryManager RepositoryManager
}

func (r *usecaseManager) ProductRepo() usecase.ProductRegistrationUseCase {
	return usecase.NewProductRegistrationUseCase(r.repositoryManager.ProductRepo())
}

func NewUseCaseManager(repositoryManager RepositoryManager) UseCaseManager {
	return &usecaseManager{repositoryManager: repositoryManager}
}
