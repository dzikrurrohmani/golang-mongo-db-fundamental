package manager

import "go-mongod/repository"

type RepositoryManager interface {
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.DbConn())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{infraManager: infraManager}
}