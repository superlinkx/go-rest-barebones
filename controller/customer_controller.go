package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/superlinkx/go-rest-barebones/app"
	"github.com/superlinkx/go-rest-barebones/decoder"
	"github.com/superlinkx/go-rest-barebones/encoder"
	"github.com/superlinkx/go-rest-barebones/sqlc"

	"github.com/gorilla/mux"
	"gopkg.in/guregu/null.v4"
)

type CustomerController struct {
	App *app.App
}

type CustomerUpdateParams struct {
	ID        int32
	FirstName null.String `json:"first_name"`
	LastName  null.String `json:"last_name"`
	Email     null.String `json:"email"`
	Phone     null.String `json:"phone"`
}

func NewCustomerController(app app.App) CustomerController {
	return CustomerController{
		App: &app,
	}
}

// CustomerPostHandler handles POST requests to /customer
func (s CustomerController) CustomerPostHandler(response http.ResponseWriter, request *http.Request) {
	var (
		customerMap map[string]string
	)

	if err := decoder.DecodeJSONRequest(request, &customerMap); err != nil {
		encoder.WriteJSONError(response, 400, fmt.Errorf("could not decode JSON: %w", err))
	} else if customer, err := customerMapToPostParams(customerMap); err != nil {
		encoder.WriteJSONError(response, 400, fmt.Errorf("validation failed: %w", err))
	} else if savedCustomer, err := s.App.Queries.CreateCustomer(s.App.Ctx, customer); err != nil {
		encoder.WriteJSONError(response, 500, fmt.Errorf("could not create customer: %w", err))
	} else {
		encoder.WriteJSONResponse(response, 200, savedCustomer)
	}
}

// CustomerGetHandler handles GET requests to /customer/{id}
func (s CustomerController) CustomerGetHandler(response http.ResponseWriter, request *http.Request) {
	if id := mux.Vars(request)["id"]; id == "" {
		encoder.WriteJSONError(response, 400, errors.New("id is required"))
	} else if numId, err := strconv.Atoi(id); err != nil {
		encoder.WriteJSONError(response, 400, fmt.Errorf("id must be a number: %w", err))
	} else if customer, err := s.App.Queries.GetCustomer(s.App.Ctx, int32(numId)); err != nil {
		encoder.WriteJSONError(response, 500, fmt.Errorf("could not get customer: %w", err))
	} else {
		encoder.WriteJSONResponse(response, 200, customer)
	}
}

// CustomerPutHandler handles PUT requests to /customer/{id}
func (s CustomerController) CustomerPutHandler(response http.ResponseWriter, request *http.Request) {
	if id := mux.Vars(request)["id"]; id == "" {
		encoder.WriteJSONError(response, 400, errors.New("id is required"))
	} else if customerParams, err := newCustomerUpdateParams(id); err != nil {
		encoder.WriteJSONError(response, 400, fmt.Errorf("invalid id: %w", err))
	} else if err := decoder.DecodeJSONRequest(request, &customerParams); err != nil {
		encoder.WriteJSONError(response, 400, fmt.Errorf("could not decode JSON: %w", err))
	} else if savedCustomer, err := s.App.Queries.UpdateCustomer(
		s.App.Ctx, convertCustomerUpdateParamsToUpdateCustomerParams(customerParams)); err != nil {
		encoder.WriteJSONError(response, 500, fmt.Errorf("could not update customer: %w", err))
	} else {
		encoder.WriteJSONResponse(response, 200, savedCustomer)
	}
}

// CustomerDeleteHandler handles DELETE requests to /customer/{id}
func (s CustomerController) CustomerDeleteHandler(response http.ResponseWriter, request *http.Request) {
	if id := mux.Vars(request)["id"]; id == "" {
		encoder.WriteJSONError(response, 400, errors.New("id is required"))
	} else if numId, err := strconv.Atoi(id); err != nil {
		encoder.WriteJSONError(response, 400, fmt.Errorf("id must be a number: %w", err))
	} else if err := s.App.Queries.DeleteCustomer(s.App.Ctx, int32(numId)); err != nil {
		encoder.WriteJSONError(response, 500, fmt.Errorf("could not delete customer: %w", err))
	} else {
		encoder.WriteJSONResponse(response, 200, nil)
	}
}

func customerMapToPostParams(customerMap map[string]string) (sqlc.CreateCustomerParams, error) {
	if customerMap["first_name"] == "" {
		return sqlc.CreateCustomerParams{}, errors.New("first_name is required")
	} else if customerMap["last_name"] == "" {
		return sqlc.CreateCustomerParams{}, errors.New("last_name is required")
	} else if customerMap["email"] == "" {
		return sqlc.CreateCustomerParams{}, errors.New("email is required")
	} else if customerMap["phone_number"] == "" {
		return sqlc.CreateCustomerParams{}, errors.New("phone_number is required")
	}

	return sqlc.CreateCustomerParams{
		FirstName: customerMap["first_name"],
		LastName:  customerMap["last_name"],
		Email:     customerMap["email"],
		Phone:     customerMap["phone_number"],
	}, nil
}

func newCustomerUpdateParams(id string) (CustomerUpdateParams, error) {
	if numId, err := strconv.Atoi(id); err != nil {
		return CustomerUpdateParams{}, fmt.Errorf("invalid id: %w", err)
	} else {
		return CustomerUpdateParams{ID: int32(numId)}, nil
	}
}

func convertCustomerUpdateParamsToUpdateCustomerParams(params CustomerUpdateParams) sqlc.UpdateCustomerParams {
	return sqlc.UpdateCustomerParams{
		ID:        params.ID,
		FirstName: params.FirstName.NullString,
		LastName:  params.LastName.NullString,
		Email:     params.Email.NullString,
		Phone:     params.Phone.NullString,
	}
}
