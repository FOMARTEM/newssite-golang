package api

import "github.com/FOMARTEM/newssite-golang/internal/entities"

type Usecase interface {
	//работа с User
	CreateUser(user entities.User) (*entities.User, error)
	CheckPasswordUser(user entities.User) (*bool, error)
	SelectUserByID(id int) (*entities.User, error)
	SelectUserByEmail(email string) (*entities.User, error)
	UpdateAdminRules(email string, admin int) (*bool, error)
	UpdateUser(user entities.User) (*entities.User, error)
	DeleteUserById(id int) error
	DeleteUserByEmail(email string) error

	//работа с Post
	CreatePost(post entities.Post) (*entities.Post, error)
	SelectPost(id int) (*entities.Post, error)
	ListPosts() ([]*entities.Post, error)
	UpdatePost(post entities.Post) (*entities.Post, error)
	DeletePost(id int) error
}
