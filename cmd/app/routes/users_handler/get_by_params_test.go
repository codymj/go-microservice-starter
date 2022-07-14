package users_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-microservice-starter/cmd/app/util"
	"go-microservice-starter/internal/users/users_dao"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetByParams_NoResponseBody(t *testing.T) {
	// init handler
	mus := mockUsersService{}
	mvs := mockValidateService{}
	services := util.Services{
		ValidatorService: &mvs,
		UserService:      &mus,
	}
	apiVersion := "/v1"
	r := mux.NewRouter()
	h := New(services)
	h.InitRoutes(r, apiVersion)

	// mock data
	queryParams := map[string]string{}

	// setup request
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", apiVersion, "users"),
		nil,
	)
	req = mux.SetURLVars(req, queryParams)

	// stub response
	expected := make([]users_dao.User, 0)
	mus.On("GetByParams", mock.Anything, queryParams).Return(expected, nil)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 204, w.Code)
}

func TestGetByParams_NoQueryParams(t *testing.T) {
	// init handler
	mus := mockUsersService{}
	mvs := mockValidateService{}
	services := util.Services{
		ValidatorService: &mvs,
		UserService:      &mus,
	}
	apiVersion := "/v1"
	r := mux.NewRouter()
	h := New(services)
	h.InitRoutes(r, apiVersion)

	// mock data
	queryParams := map[string]string{}

	// setup request
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", apiVersion, "users"),
		nil,
	)
	req = mux.SetURLVars(req, queryParams)

	// stub response
	expectedJsn := []byte(`
		[
			{
				"id": "489079b4-a254-4e91-b6ad-1072db48d264",
				"username": "user0",
				"email": "user0@gmail.com",
				"isVerified": false,
				"createdOn": 1657395460618,
				"updatedOn": 1657396010853
			},
			{
				"id": "489079b4-a254-4e91-b6ad-1072db48d264",
				"username": "user1",
				"email": "user1@gmail.com",
				"isVerified": false,
				"createdOn": 1657395460790,
				"updatedOn": 1657396012638
			}
		]
		`)
	expected := make([]users_dao.User, 2)
	_ = json.Unmarshal(expectedJsn, &expected)
	mus.On("GetByParams", mock.Anything, queryParams).Return(expected, nil)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 200, w.Code)

	expectedCompact := new(bytes.Buffer)
	_ = json.Compact(expectedCompact, expectedJsn)
	assert.Equal(t, expectedCompact.String(), w.Body.String())
}
