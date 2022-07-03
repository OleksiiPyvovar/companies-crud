package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

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
}

func (a *API) CompanyDeleteByIDHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.handleError(w, fmt.Errorf("bad id format '%s': %v", params.ByName("id"), err), http.StatusBadRequest)
		return
	}

	ok := a.CompaniesService.DeleteByID(id)
	_ = encodeResponse(w, apiv1.DeleteCompayResponse{Success: ok})
}

func (a *API) CompanyGetByIDHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.handleError(w, fmt.Errorf("bad id format '%s': %v", params.ByName("id"), err), http.StatusBadRequest)
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
	filter := &domain.ListFilter{
		Limit:      a.getListLimitFromValuesOrDefaults(r.URL.Query()),
		Attributes: getAttrFilterFromValues(r.URL.Query()),
	}

	companies, err := a.CompaniesService.List(filter)
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

func (a *API) getListLimitFromValuesOrDefaults(values url.Values) int {
	var (
		limit int
		err   error
	)

	limit, err = strconv.Atoi(values.Get("limit"))
	if err != nil {
		limit = a.Config.DefaultListLimit
	}

	return limit
}

func getAttrFilterFromValues(values url.Values) *domain.Company {
	return &domain.Company{
		Name:    values.Get("name"),
		Code:    values.Get("code"),
		Country: values.Get("country"),
		Website: values.Get("website"),
		Phone:   values.Get("phone"),
	}
}

func encodeResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(resp)
}
