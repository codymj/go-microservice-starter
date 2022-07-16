package users_handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-microservice-starter/cmd/app/util"
	"go-microservice-starter/internal/users/users_dao"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetById_Success(t *testing.T) {
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
	idParam := uuid.New()
	pathParam := map[string]string{
		"id": idParam.String(),
	}

	// setup request
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam.String()),
		nil,
	)
	req = mux.SetURLVars(req, pathParam)

	// stub response
	expected := users_dao.User{
		Id:         uuid.New(),
		Username:   "test",
		Email:      "test@gmail.com",
		IsVerified: false,
		CreatedOn:  1657384667179,
		UpdatedOn:  1657384667179,
	}
	mus.On("GetById", mock.Anything, idParam).Return(expected, nil)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 200, w.Code)
}

func TestGetById_InvalidPathParam(t *testing.T) {
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
	idParam := "1"
	pathParam := map[string]string{
		"id": idParam,
	}

	// setup request
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam),
		nil,
	)
	req = mux.SetURLVars(req, pathParam)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 400, w.Code)
}

func TestGetById_InternalError(t *testing.T) {
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
	idParam := uuid.New()
	pathParam := map[string]string{
		"id": idParam.String(),
	}

	// setup request
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam.String()),
		nil,
	)
	req = mux.SetURLVars(req, pathParam)

	// stub response
	expected := users_dao.User{}
	err := errors.New("internal error")
	mus.On("GetById", mock.Anything, idParam).Return(expected, err)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 500, w.Code)
}

func TestGetById_NoUserFound(t *testing.T) {
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
	idParam := uuid.New()
	pathParam := map[string]string{
		"id": idParam.String(),
	}

	// setup request
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam.String()),
		nil,
	)
	req = mux.SetURLVars(req, pathParam)

	// stub response
	expected := users_dao.User{
		Id:         uuid.Nil,
		Username:   "test",
		Email:      "test@gmail.com",
		IsVerified: false,
		CreatedOn:  1657384667179,
		UpdatedOn:  1657384667179,
	}
	mus.On("GetById", mock.Anything, idParam).Return(expected, nil)

	// invocation
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, 204, w.Code)
}
