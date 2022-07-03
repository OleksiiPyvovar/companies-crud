package v1

type CompaniesListResponse struct {
	// Company items
	//
	// required: true
	Items []CompanyResponse `json:"items"`
}

type CompanyResponse struct {
	// Company ID auto generated for each company
	//
	// required: true
	// example: 1
	ID int `json:"id"`
	// Name
	//
	// required: true
	// example: EPAM
	Name string `json:"name"`
	// Code
	//
	// required: true
	// example: 00101
	Code string `json:"code"`
	// Country
	//
	// required: true
	// example: Ukraine
	Country string `json:"country"`
	// Website
	//
	// required: false
	// example: https://www.epam.com/
	Website string `json:"website"`
	// Phone
	//
	// required: false
	// example: +38063334422
	Phone string `json:"phone"`
}

type CompanyRequest struct {
	// Name
	//
	// required: true
	// example: EPAM
	Name string `json:"name"`
	// Code
	//
	// required: true
	// example: 00101
	Code string `json:"code"`
	// Country
	//
	// required: true
	// example: Ukraine
	Country string `json:"country"`
	// Website
	//
	// required: false
	// example: https://www.epam.com/
	Website string `json:"website"`
	// Phone
	//
	// required: false
	// example: +38063334422
	Phone string `json:"phone"`
}

type DeleteCompayResponse struct {
	// Success
	//
	// required: false
	// example: true
	Success bool
}
