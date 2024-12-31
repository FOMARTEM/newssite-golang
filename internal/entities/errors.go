package entities

import error

var (
	//ошибки постов
	ErrPostNotFound = error.new("post not found")
	
	ErrPostNameConflict = error.new("post name conflict")

	//ошибки пользователя
	ErrUserNotFound = error.new("user not found")

	ErrUserAlreadyExist = error.new("user already exists")
	
	ErrUserNotAdmin	= error.new("user not admin")
	
	ErrUserNameConflict = error.new("user name conflict")
	ErrUserEmailConflict = error.new("user email conflict")
	ErrUserPasswordConflict = error.new("user password conflict")
)