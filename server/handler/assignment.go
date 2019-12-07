package handler

import (
	"net/http"

	"kilvish.io/cerebro/entity"
	"kilvish.io/cerebro/service"
)

type AutoAssignHandler struct {
	bs service.IBucketingService
}

func NewAutoAssignHandler() *AutoAssignHandler {
	return &AutoAssignHandler{
		bs: &service.BucketingService{},
	}
}

func deserializeToAutoAssignRequestAndValidate(r *http.Request) (entity.AutoAssignRequest, error) {
	return entity.AutoAssignRequest{}, nil
}

func (a *AutoAssignHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	autoAssignRequest, err := deserializeToAutoAssignRequestAndValidate(r)
	if err != nil {
		// Do something and send appropriate response
	}
	go a.bs.Assign(&autoAssignRequest)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("done"))
}
