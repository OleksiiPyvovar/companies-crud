package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *API) MiddlewareAuthentication(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := a.Authorizer.Validate(r)
		if err != nil {
			a.handleError(w, errors.New("Unauthorized"), http.StatusUnauthorized)
			return
		}
		next(w, r, params)
	}
}
