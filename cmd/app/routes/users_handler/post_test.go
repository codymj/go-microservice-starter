package users_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-microservice-starter/cmd/app/util"
	"go-microservice-starter/internal/users"
	"go-microservice-starter/internal/users/users_dao"
	"go-microservice-starter/internal/validate"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost_Success(t *testing.T) {
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
	body := []byte(`
	{
		"username": "new_user",
		"password": "s3cr3t",
		"email": "new_user@yahoo.com"
	}
	`)
	var postUserReq users.PostUsersRequest
	_ = json.Unmarshal(body, &postUserReq)

	// setup request
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", apiVersion, "users"),
		bytes.NewReader(body),
	)

	// stub response
	expectedJsn := []byte(`
	{
		"id": "489079b4-a254-4e91-b6ad-1072db48d264",
		"username": "new_user",
		"email": "new_user@yahoo.com",
		"isVerified": false,
		"createdOn": 1657395460618,
		"updatedOn": 1657395460618
	}
	`)
	var expected users_dao.User
	_ = json.Unmarshal(expectedJsn, &expected)
	mvs.On("ValidatePostUsers", mock.Anything, body).Return([]string{}, nil)
	mus.On("Save", mock.Anything, postUserReq).Return(expected, nil)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 201, w.Code)
}

func TestPost_ValidationError(t *testing.T) {
	// init handler
	mus := mockUsersService{}
	vs := validate.New()
	services := util.Services{
		ValidatorService: vs,
		UserService:      &mus,
	}
	apiVersion := "/v1"
	r := mux.NewRouter()
	h := New(services)
	h.InitRoutes(r, apiVersion)

	// mock data
	body := []byte(`
	{
		"username": "new_user"
		"email": "new_user@yahoo.com"
	}
	`)
	var postUserReq users.PostUsersRequest
	_ = json.Unmarshal(body, &postUserReq)

	// setup request
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", apiVersion, "users"),
		bytes.NewReader(body),
	)

	// stub response
	expectedJsn := []byte(`
	{
		"id": "489079b4-a254-4e91-b6ad-1072db48d264",
		"username": "new_user",
		"email": "new_user@yahoo.com",
		"isVerified": false,
		"createdOn": 1657395460618,
		"updatedOn": 1657395460618
	}
	`)
	var expected users_dao.User
	_ = json.Unmarshal(expectedJsn, &expected)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 400, w.Code)
}

func TestPost_InternalError(t *testing.T) {
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
	body := []byte(`
	{
		"username": "new_user",
		"password": "s3cr3t",
		"email": "new_user@yahoo.com"
	}
	`)
	var postUserReq users.PostUsersRequest
	_ = json.Unmarshal(body, &postUserReq)

	// setup request
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", apiVersion, "users"),
		bytes.NewReader(body),
	)

	// stub response
	err := errors.New("internal error")
	mvs.On("ValidatePostUsers", mock.Anything, body).Return([]string{}, nil)
	mus.On("Save", mock.Anything, postUserReq).Return(users_dao.User{}, err)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 500, w.Code)
}
