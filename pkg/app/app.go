package app

import (
	"github.com/OleksiiPyvovar/companies-crud/pkg/domain"
)

type Service interface {
	Create(company *domain.Company) error
	Update(company *domain.Company) error
	DeleteByID(id int) bool
	GetByID(id int) (domain.Company, error)
	List(options *domain.ListFilter) ([]domain.Company, error)
}

type service struct {
	repository domain.Repository
}

func NewCompaniesService(repo domain.Repository) Service {
	return &service{
		repository: repo,
	}
}

func (svc *service) Create(company *domain.Company) error {
	return svc.repository.Create(company)
}

func (svc *service) Update(company *domain.Company) error {
	return svc.repository.Update(company)
}

func (svc *service) DeleteByID(id int) bool {
	return svc.repository.DeleteByID(id)
}

func (svc *service) GetByID(id int) (domain.Company, error) {
	return svc.repository.GetByID(id)
}

func (svc *service) List(options *domain.ListFilter) ([]domain.Company, error) {
	return svc.repository.List(options)
}
