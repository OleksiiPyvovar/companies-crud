package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (a *API) handleError(w http.ResponseWriter, err error, code int) {
	a.Logger.Infof("error response %d: %v", code, err)

	err = json.NewEncoder(w).Encode(errorResponse{Code: code, Message: err.Error()})
	if err != nil {
		a.Logger.Errorf("handle error: failed to encode error: %v", err)
	}
}
