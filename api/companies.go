package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	apiv1 "github.com/OleksiiPyvovar/companies-crud/api/v1"
	"github.com/OleksiiPyvovar/companies-crud/pkg/domain"

	"github.com/julienschmidt/httprouter"
)

func (a *API) CompanyCreateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var company domain.Company

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		a.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = a.CompaniesService.Create(&company)
	if err != nil {
		a.handleError(w, err, http.StatusInternalServerError)
		return
	}

	_ = encodeResponse(w, apiv1.CompanyResponse{
		ID:      company.ID,
		Name:    company.Name,
		Code:    company.Code,
		Country: company.Country,
		Website: company.Website,
		Phone:   company.Phone,
	})
}

func (a *API) CompanyUpdateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var company domain.Company

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		a.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = a.CompaniesService.Update(&company)
	if err != nil {
		a.handleError(w, err, http.StatusInternalServerError)
		return
	}

	_ = encodeResponse(w, apiv1.CompanyResponse{
		ID:      company.ID,
		Name:    company.Name,
		Code:    company.Code,
		Country: company.Country,
		Website: company.Website,
		Phone:   company.Phone,
	})
}

func (a *API) CompanyDeleteByIDHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.handleError(w, fmt.Errorf("bad id format '%s': %w", params.ByName("id"), err), http.StatusBadRequest)
		return
	}

	ok := a.CompaniesService.DeleteByID(id)
	_ = encodeResponse(w, apiv1.DeleteCompayResponse{Success: ok})
}

func (a *API) CompanyGetByIDHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.handleError(w, fmt.Errorf("bad id format '%s': %w", params.ByName("id"), err), http.StatusBadRequest)
		return
	}

	company, err := a.CompaniesService.GetByID(id)
	if err != nil {
		a.handleError(w, err, http.StatusInternalServerError)
		return
	}

	_ = encodeResponse(w, apiv1.CompanyResponse{
		ID:      company.ID,
		Name:    company.Name,
		Code:    company.Code,
		Country: company.Country,
		Website: company.Website,
		Phone:   company.Phone,
	})
}

func (a *API) CompanyListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	companyQueryAttr := []string{"name", "code", "country", "website", "phone"}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = a.Config.DefaultListLimit
	}

	var filters []domain.Filter
	for _, attr := range companyQueryAttr {
		if value := r.URL.Query().Get(attr); value != "" {
			val, operator := splitValueAndOperator(value)
			filters = append(filters, domain.Filter{
				Attr:     attr,
				Value:    val,
				Operator: operator,
			})
		}
	}

	companies, err := a.CompaniesService.List(domain.ListFilter{
		Limit:   limit,
		Filters: filters,
	})
	if err != nil {
		a.handleError(w, err, http.StatusInternalServerError)
		return
	}

	resp := apiv1.CompaniesListResponse{
		Items: make([]apiv1.CompanyResponse, 0, len(companies)),
	}

	for i := range companies {
		resp.Items = append(resp.Items, apiv1.CompanyResponse{
			ID:      companies[i].ID,
			Name:    companies[i].Name,
			Code:    companies[i].Code,
			Country: companies[i].Country,
			Website: companies[i].Website,
			Phone:   companies[i].Phone,
		})
	}

	_ = encodeResponse(w, resp)
}

func encodeResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(resp)
}

func splitValueAndOperator(v string) (string, string) {
	parts := strings.Split(v, "|")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}

	return v, ""
}
