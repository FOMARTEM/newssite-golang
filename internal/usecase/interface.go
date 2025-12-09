package usecase

import (
  "github.com/FOMARTEM/newssite-golang/internal/entities"
)

type Provider interface {
  //работа с post
  InsertPost(post entities.Post) (*entities.Post, error)

  SelectPostById(id int) (*entities.Post, error)
  SelectAllPosts(limit int, offset int) ([]*entities.Post, error)
  SelectUserPosts(userId int, limit int, offset int) ([]*entities.Post, error)

  UpdatePostById(post entities.Post) (*entities.Post, error)
  EditVisibilityById(id int, hide int, UpdateDate string) error

  DeletePostById(id int) error

  //работа с user
  InsertUser(user entities.User) (*entities.User, error)

  SelectUserByEmail(email string) (*entities.User, error)
  SelectUserById(id int) (*entities.User, error)

  SelectUserRulesByEmail(email string) (*int, error)
  SelectUserRulesById(id int) (*int, error)

  SelectUserPasswordByEmail(email string) (*string, error)

  UpdateUserById(user entities.User) (*entities.User, error)
  UpdateUserAdminRulesByEmail(email string, admin int) error

  DeleteUserById(id int) error
  DeleteUserByEmail(email string) error

  //работа сomments
  InsertComment(comment entities.Comment) (*entities.Comment, error)

  GetCommentsForPost(post_id int, limit int, offset int) ([]*entities.Comment, error)

  UpdateComment(comment entities.Comment) (*entities.Comment, error)

  DeleteComment(id int) error

  DeleteCommentsInPost(postId int) error
}
