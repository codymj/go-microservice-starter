package users_handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"go-microservice-starter/internal/users"
	"go-microservice-starter/internal/users/users_dao"
)

/*
 * mocked validation service
 */

type mockValidateService struct {
	mock.Mock
}

func (m *mockValidateService) ValidatePostUsers(ctx context.Context, body []byte) ([]string, error) {
	args := m.Called(ctx, body)
	return args.Get(0).([]string), args.Error(1)
}

func (m *mockValidateService) ValidatePutUsersId(ctx context.Context, body []byte) ([]string, error) {
	args := m.Called(ctx, body)
	return args.Get(0).([]string), args.Error(1)
}

/*
 * mocked user service
 */

type mockUsersService struct {
	mock.Mock
}

func (m *mockUsersService) DeleteById(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockUsersService) GetByParams(ctx context.Context, params map[string]string) ([]users_dao.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]users_dao.User), args.Error(1)
}

func (m *mockUsersService) GetById(ctx context.Context, id uuid.UUID) (users_dao.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(users_dao.User), args.Error(1)
}

func (m *mockUsersService) Save(ctx context.Context, r users.PostUsersRequest) (users_dao.User, error) {
	args := m.Called(ctx, r)
	return args.Get(0).(users_dao.User), args.Error(1)
}

func (m *mockUsersService) UpdateById(ctx context.Context, id uuid.UUID, r users.PutUsersIdRequest) (users_dao.User, error) {
	args := m.Called(ctx, id, r)
	return args.Get(0).(users_dao.User), args.Error(1)
}
