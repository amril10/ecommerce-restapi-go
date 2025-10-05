package service_test

import (
	"errors"
	"testing"

	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/model"
	"github.com/amril10/rest-api-go/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) EmailExist(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

func (m *MockAuthRepository) Register(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAuthRepository) GetUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := service.NewAuthService(mockRepo)

	req := &dto.RegisterRequest{
		Name:                 "John Doe",
		Email:                "john@gmail.com",
		Password:             "contoh123",
		PasswordConfirmation: "contoh123",
		RoleID:               1,
	}

	mockRepo.On("EmailExist", req.Email).Return(false)
	mockRepo.On("Register", mock.AnythingOfType("*model.User")).Return(nil)

	err := service.Register(req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	authService := service.NewAuthService(mockRepo)

	req := &dto.RegisterRequest{
		Name:                 "John Doe",
		Email:                "john@gmail.com",
		Password:             "contoh123",
		PasswordConfirmation: "contoh123",
		RoleID:               1,
	}

	// Mock: email sudah digunakan
	mockRepo.On("EmailExist", req.Email).Return(true)

	err := authService.Register(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email")
	mockRepo.AssertExpectations(t)
}

func TestRegister_SaveUserError(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	authService := service.NewAuthService(mockRepo)

	req := &dto.RegisterRequest{
		Name:                 "John Doe",
		Email:                "john@gmail.com",
		Password:             "contoh123",
		PasswordConfirmation: "contoh123",
		RoleID:               1,
	}

	mockRepo.On("EmailExist", req.Email).Return(false)
	mockRepo.On("Register", mock.AnythingOfType("*model.User")).Return(errors.New("failed to insert user"))

	err := authService.Register(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to insert user")
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := service.NewAuthService(mockRepo)

	req := &dto.LoginRequest{
		Email:    "john@gmail.com",
		Password: "contoh123",
	}

	user := &model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@gmail.com",
		Password: "contoh123",
	}

	mockRepo.On("GetUserByEmail", req.Email).Return(user, nil)

	resp, err := service.Login(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, user.ID, resp.ID)
	assert.Equal(t, user.Name, resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := service.NewAuthService(mockRepo)

	req := &dto.LoginRequest{
		Email:    "notfound@gmail.com",
		Password: "123456",
	}

	mockRepo.On("GetUserByEmail", req.Email).Return(nil, errors.New("user not found"))

	resp, err := service.Login(req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid credentials")

	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := service.NewAuthService(mockRepo)

	req := &dto.LoginRequest{
		Email:    "notfound@gmail.com",
		Password: "123456",
	}

	user := &model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@gmail.com",
		Password: "correctpass",
	}

	mockRepo.On("GetUserByEmail", req.Email).Return(user, nil)

	resp, err := service.Login(req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid credentials")

	mockRepo.AssertExpectations(t)
}
