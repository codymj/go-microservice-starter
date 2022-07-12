package users_handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"go-microservice-starter/cmd/app/routes/users_handler"
	"go-microservice-starter/cmd/app/util"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("DeleteById()", func() {
	// stub
	mus := mockUsersService{}
	mvs := mockValidateService{}
	req := &http.Request{}
	w := &httptest.ResponseRecorder{}
	apiVersion := "/v1"
	r := &mux.Router{}

	BeforeEach(func() {
		// need to reset recorder before each test or same code gets returned
		w = httptest.NewRecorder()

		// initialize handler
		r = mux.NewRouter()
		services := util.Services{
			ValidatorService: &mvs,
			UserService:      &mus,
		}
		h := users_handler.New(services)
		h.InitRoutes(r, apiVersion)
	})

	It("should successfully delete a user", func() {
		// mock data
		idParam := uuid.New()
		req, _ = http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam.String()),
			nil,
		)
		pathParam := map[string]string{
			"id": idParam.String(),
		}
		req = mux.SetURLVars(req, pathParam)

		// stub response
		mus.On("DeleteById", mock.Anything, idParam).Return(nil)

		// invocation
		r.ServeHTTP(w, req)

		// assertions
		Expect(w.Code).To(Equal(200))
	})

	It("should return 400 for bad path param", func() {
		// mock data
		idParam := "1"
		req, _ = http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam),
			nil,
		)
		pathParam := map[string]string{
			"id": idParam,
		}
		req = mux.SetURLVars(req, pathParam)

		// invocation
		r.ServeHTTP(w, req)

		// assertions
		Expect(w.Code).To(Equal(400))
	})

	It("should return 500 for internal error", func() {
		// mock data
		idParam := uuid.New()
		req, _ = http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/%s/%s", apiVersion, "users", idParam.String()),
			nil,
		)
		pathParam := map[string]string{
			"id": idParam.String(),
		}
		req = mux.SetURLVars(req, pathParam)

		// stub response
		err := errors.New("internal error")
		mus.On("DeleteById", mock.Anything, idParam).Return(err)

		// invocation
		r.ServeHTTP(w, req)

		// assertions
		Expect(w.Code).To(Equal(500))
	})
})
