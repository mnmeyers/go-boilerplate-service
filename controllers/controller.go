package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go-boilerplate/models"
	"go-boilerplate/services"
	httpResponse "go-boilerplate/utilities/http"
	"net/http"
)

// Controller defines the interface for working with payments
type Controller interface {
	List(writer http.ResponseWriter, request *http.Request)
	Post(writer http.ResponseWriter, request *http.Request)
}

type ControllerImpl struct {
	service services.Service
}

var _ Controller = (*ControllerImpl)(nil)

// RenderJSON is an alias of method to render JSON for easy mocking in tests
var RenderJSON = render.JSON

func GetController() Controller {
	return &ControllerImpl{
		service: services.GetService(),
	}
}

func (controller *ControllerImpl) List(w http.ResponseWriter, req *http.Request) {
	accountId := getParam(req, "account_id")
	if accountId == "" {
		w.WriteHeader(http.StatusBadRequest)
		RenderJSON(w, req, httpResponse.BadRequest("Invalid URL parameter \"account_id\" received"))
		return
	}

	payments, err := controller.service.Get(req.Context(), accountId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		RenderJSON(w, req, httpResponse.BadRequest(err.Error()))
		return
	}

	if payments == nil {
		httpResponse.RenderJSON(w, req, []models.DTO{})
		return
	}

	httpResponse.RenderJSON(w, req, payments)
}

func (controller *ControllerImpl) Post(w http.ResponseWriter, req *http.Request) {
	paymentReq, err := getPostBody(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		RenderJSON(w, req, httpResponse.BadRequest(err.Error()))
		return
	}

	if paymentReq.CustomerId == "" || paymentReq.Amount == "" {
		w.WriteHeader(http.StatusBadRequest)
		RenderJSON(w, req, httpResponse.BadRequest("\"account_id\" and \"amount\" are required fields in the POST body"))
		return
	}

	createdPayment, err := controller.service.Create(req.Context(), *paymentReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		RenderJSON(w, req, httpResponse.BadRequest(err.Error()))
		return
	}

	httpResponse.RenderJSON(w, req, createdPayment)
}

// getPostBody wraps operations for extracting request body for easy mocking in tests
var getPostBody = func(r *http.Request) (*models.PostBody, error) {
	var data models.PostBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	return &data, err
}

// getParam wraps chi.URLParam for easy mocking in test
var getParam = func(r *http.Request, p string) string {
	return chi.URLParam(r, p)
}
