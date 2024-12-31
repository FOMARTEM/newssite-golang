package entities

import "errors"

var (
	//ошибки постов
	ErrPostNotFound = errors.New("post not found")

	ErrPostNameConflict = errors.New("post name conflict")

	//ошибки пользователя
	ErrUserNotFound = errors.New("user not found")

	ErrUserAlreadyExist = errors.New("user already exists")

	ErrUserNotAdmin = errors.New("user not admin")

	ErrUserNameConflict     = errors.New("user name conflict")
	ErrUserEmailConflict    = errors.New("user email conflict")
	ErrUserPasswordConflict = errors.New("user password conflict")
)
