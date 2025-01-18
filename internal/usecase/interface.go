package usecase

import "github.com/FOMARTEM/newssite-golang/internal/entities"

type Provider interface {
	//работа с post
	InsertPost(post entities.Post) (*entities.Post, error)

	SelectPostById(id int) (*entities.Post, error)
	SelectAllPosts() ([]*entities.Post, error)

	UpdatePostById(post entities.Post) (*entities.Post, error)

	DeletePostById(id int) error

	//работа с user
	InsertUser(user entities.User) (*entities.User, error)

	SelectUserById(id int) (*entities.User, error)
	SelectUserByEmail(email string) (*entities.User, error)
	SelectUserPasswordByEmail(email string) (*string, error)

	UpdateUserById(user entities.User) (*entities.User, error)
	UpdateUserAdminRulesById(id int, admin int) error
	UpdateUserAdminRulesByEmail(email string, admin int) error

	CheckUserIsAdminById(id int) (*int, error)
	CheckUserIsAdminByEmail(email string) (*int, error)

	DeleteUserById(id int) error
	DeleteUserByEmail(email string) error
}
