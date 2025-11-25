package usecase_test

import (
	"errors"
	"testing"

	"github.com/FOMARTEM/newssite-golang/internal/provider/mocks"
	"github.com/FOMARTEM/newssite-golang/internal/usecase"

	"github.com/FOMARTEM/newssite-golang/internal/entities"

	"github.com/stretchr/testify/require"
)

//
// CREATE USER
//

func TestCreateUser_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	input := entities.User{
		Name:     "John",
		Email:    "john@mail.com",
		Password: "12345678",
	}

	created := &entities.User{ID: 1, Name: input.Name, Email: input.Email}

	mockP.On("SelectUserByEmail", input.Email).Return(nil, nil)
	mockP.On("InsertUser", input).Return(created, nil)

	res, err := u.CreateUser(input)
	require.NoError(t, err)
	require.Equal(t, created, res)

	mockP.AssertExpectations(t)
}

func TestCreateUser_EmailConflict(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	existing := &entities.User{ID: 1, Email: "a@a.com"}

	mockP.On("SelectUserByEmail", "a@a.com").Return(existing, nil)

	res, err := u.CreateUser(entities.User{Email: "a@a.com"})
	require.Nil(t, res)
	require.ErrorIs(t, err, entities.ErrUserEmailConflict)

	mockP.AssertExpectations(t)
}

func TestCreateUser_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("SelectUserByEmail", "test@test.com").
		Return(nil, errors.New("db error"))

	res, err := u.CreateUser(entities.User{Email: "test@test.com"})
	require.Nil(t, res)
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}

//
// CHECK PASSWORD USER
//

func TestCheckPasswordUser_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	input := entities.User{Email: "a@a.com", Password: "pass123"}
	pass := "pass123"

	mockP.On("SelectUserPasswordByEmail", input.Email).Return(&pass, nil)

	result, err := u.CheckPasswordUser(input)
	require.NoError(t, err)
	require.True(t, *result)

	mockP.AssertExpectations(t)
}

func TestCheckPasswordUser_WrongPassword(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	input := entities.User{Email: "a@a.com", Password: "wrong"}
	pass := "correct"

	mockP.On("SelectUserPasswordByEmail", input.Email).Return(&pass, nil)

	result, err := u.CheckPasswordUser(input)
	require.False(t, *result)
	require.ErrorIs(t, err, entities.ErrUserLoginConflict)

	mockP.AssertExpectations(t)
}

func TestCheckPasswordUser_UserNotFound(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("SelectUserPasswordByEmail", "a@a.com").Return(nil, nil)

	result, err := u.CheckPasswordUser(entities.User{Email: "a@a.com"})
	require.False(t, *result)
	require.ErrorIs(t, err, entities.ErrUserNotFound)

	mockP.AssertExpectations(t)
}

//
// SELECT USER BY ID
//

func TestSelectUserByID_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	user := &entities.User{ID: 1, Email: "a@a.com"}

	mockP.On("SelectUserById", 1).Return(user, nil)

	res, err := u.SelectUserByID(1)
	require.NoError(t, err)
	require.Equal(t, user, res)

	mockP.AssertExpectations(t)
}

func TestSelectUserByID_NotFound(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("SelectUserById", 1).Return(nil, nil)

	res, err := u.SelectUserByID(1)
	require.Nil(t, res)
	require.ErrorIs(t, err, entities.ErrUserNotFound)

	mockP.AssertExpectations(t)
}

func TestSelectUserByID_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("SelectUserById", 1).Return(nil, errors.New("db error"))

	res, err := u.SelectUserByID(1)
	require.Nil(t, res)
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}

//
// SELECT USER BY EMAIL
//

func TestSelectUserByEmail_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	user := &entities.User{Email: "a@a.com"}

	mockP.On("SelectUserByEmail", "a@a.com").Return(user, nil)

	res, err := u.SelectUserByEmail("a@a.com")
	require.NoError(t, err)
	require.Equal(t, user, res)

	mockP.AssertExpectations(t)
}

func TestSelectUserByEmail_NotFound(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("SelectUserByEmail", "a@a.com").Return(nil, nil)

	res, err := u.SelectUserByEmail("a@a.com")
	require.Nil(t, res)
	require.ErrorIs(t, err, entities.ErrUserNotFound)

	mockP.AssertExpectations(t)
}

func TestSelectUserByEmail_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("SelectUserByEmail", "a@a.com").Return(nil, errors.New("db error"))

	res, err := u.SelectUserByEmail("a@a.com")
	require.Nil(t, res)
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}

//
// UPDATE ADMIN RULES
//

func TestUpdateAdminRules_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("UpdateUserAdminRulesByEmail", "a@a.com", 1).Return(nil)

	res, err := u.UpdateAdminRules("a@a.com", 1)
	require.NoError(t, err)
	require.True(t, *res)

	mockP.AssertExpectations(t)
}

func TestUpdateAdminRules_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("UpdateUserAdminRulesByEmail", "a@a.com", 1).
		Return(errors.New("db error"))

	res, err := u.UpdateAdminRules("a@a.com", 1)
	require.False(t, *res)
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}

//
// UPDATE USER
//

func TestUpdateUser_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	input := entities.User{ID: 1, Name: "John"}
	updated := &entities.User{ID: 1, Name: "John Updated"}

	mockP.On("UpdateUserById", input).Return(updated, nil)

	res, err := u.UpdateUser(input)
	require.NoError(t, err)
	require.Equal(t, updated, res)

	mockP.AssertExpectations(t)
}

func TestUpdateUser_NotFound(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	input := entities.User{ID: 1}

	mockP.On("UpdateUserById", input).Return(nil, nil)

	res, err := u.UpdateUser(input)
	require.Nil(t, res)
	require.ErrorIs(t, err, entities.ErrUserNotFound)

	mockP.AssertExpectations(t)
}

func TestUpdateUser_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	input := entities.User{ID: 1}

	mockP.On("UpdateUserById", input).Return(nil, errors.New("db error"))

	res, err := u.UpdateUser(input)
	require.Nil(t, res)
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}

//
// DELETE USER BY ID
//

func TestDeleteUserById_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("DeleteUserById", 1).Return(nil)

	err := u.DeleteUserById(1)
	require.NoError(t, err)

	mockP.AssertExpectations(t)
}

func TestDeleteUserById_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("DeleteUserById", 1).Return(errors.New("db error"))

	err := u.DeleteUserById(1)
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}

//
// DELETE USER BY EMAIL
//

func TestDeleteUserByEmail_Success(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("DeleteUserByEmail", "a@a.com").Return(nil)

	err := u.DeleteUserByEmail("a@a.com")
	require.NoError(t, err)

	mockP.AssertExpectations(t)
}

func TestDeleteUserByEmail_DBError(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	mockP.On("DeleteUserByEmail", "a@a.com").Return(errors.New("db error"))

	err := u.DeleteUserByEmail("a@a.com")
	require.EqualError(t, err, "db error")

	mockP.AssertExpectations(t)
}
