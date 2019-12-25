package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	if r.Body == nil {
		return entity.AutoAssignRequest{}, errors.New("EmptyBody")
	}
	defer r.Body.Close()

	request := &entity.AutoAssignRequest{}
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, request)

	if err != nil {
		return entity.AutoAssignRequest{}, err
	}
	return *request, nil
}

func (a *AutoAssignHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	autoAssignRequest, err := deserializeToAutoAssignRequestAndValidate(r)
	if err != nil {
		// Do something and send appropriate response
	}
	fmt.Println(autoAssignRequest)
	go a.bs.Assign(&autoAssignRequest)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"done"}`))
}
