package domain

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type ListFilter struct {
	Limit      int
	Attributes *Company
}

type Repository interface {
	Create(company *Company) error
	Update(company *Company) error
	DeleteByID(id int) bool
	GetByID(id int) (Company, error)
	List(options *ListFilter) ([]Company, error)
}
