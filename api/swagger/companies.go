package swagger

import (
	apiv1 "github.com/OleksiiPyvovar/companies-crud/api/v1"
)

// swagger:route GET /api/v1/companies getCompanies
//
// Returns list of Companies
//
// Responses:
// 200: CompaniesListRespDef

// JSON parameters
// swagger:parameters getCompanies
type CompaniesListRequestDef struct {
	// The numbers of items to return
	//
	// in: query
	Limit int `json:"limit,omitempty"`
	// Filtring parameter
	//
	// in: query
	Name string `json:"name,omitempty"`
	// Filtring parameter
	//
	// in: query
	Code string `json:"code,omitempty"`
	// Filtring parameter
	//
	// in: query
	Country string `json:"country,omitempty"`
	// Filtring parameter
	//
	// in: query
	Website string `json:"website,omitempty"`
	// Filtring parameter
	//
	// in: query
	Phone string `json:"phone,omitempty"`
}

// JSON Response
// swagger:response
type CompaniesListRespDef struct {
	// in: body
	Body apiv1.CompaniesListResponse
}

// swagger:route GET /api/v1/companies/{id} getCompany
//
// Returns company by ID
//
// Responses:
// 200: CompanyRespDef

// JSON parameters
// swagger:parameters getCompany
type CompanyRequestDef struct {
	// Product ID
	//
	// required: true
	// in: path
	ID int `json:"id"`
}

// JSON Response
// swagger:response
type CompanyRespDef struct {
	// in: body
	Body apiv1.CompanyResponse
}

// swagger:route POST /api/v1/companies/ createCompanies
//
// Create a company
//
// Responses:
// 200: CompanyRespDef
// Security:
// - bearer: []

// JSON parameters
// swagger:parameters createCompanies
type CompanyCreateRequestDef struct {
	// in: body
	Body apiv1.CompanyRequest
}

// swagger:route DELETE /api/v1/companies/{id} deleteCompany
//
// Delete company by ID
//
// Responses:
// 200: CompanyDeleteRespDef
// Security:
// - bearer: []

// JSON parameters
// swagger:parameters deleteCompany
type CompanyDeleteRequestDef struct {
	// Product ID
	//
	// required: true
	// in: path
	ID int `json:"id"`
}

// JSON Response
// swagger:response
type CompanyDeleteRespDef struct {
	// in: body
	Body apiv1.DeleteCompayResponse
}

// swagger:route PUT /api/v1/companies/ updateCompany
//
// Update company
//
// Responses:
// 200: CompanyRespDef
// Security:
// - bearer: []

// JSON parameters
// swagger:parameters updateCompany
type UpdateCompanyRequestDef struct {
	// in: body
	Body apiv1.CompanyResponse
}
