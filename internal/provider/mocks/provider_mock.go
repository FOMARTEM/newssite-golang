package mocks

import (
  "github.com/stretchr/testify/mock"

  "github.com/FOMARTEM/newssite-golang/internal/entities"
)

// ProviderMock полностью реализует интерфейс usecase.Provider
type ProviderMock struct {
  mock.Mock
}

// ------------ POST ------------

func (m *ProviderMock) InsertPost(post entities.Post) (*entities.Post, error) {
  args := m.Called(post)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *ProviderMock) SelectPostById(id int) (*entities.Post, error) {
  args := m.Called(id)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *ProviderMock) SelectAllPosts(limit int, offset int) ([]*entities.Post, error) {
  args := m.Called()
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).([]*entities.Post), args.Error(1)
}

func (m *ProviderMock) SelectUserPosts(userId int, limit int, offset int) ([]*entities.Post, error) {
  args := m.Called(userId)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).([]*entities.Post), args.Error(1)
}

func (m *ProviderMock) UpdatePostById(post entities.Post) (*entities.Post, error) {
  args := m.Called(post)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *ProviderMock) EditVisibilityById(id int, hide int, updateDate string) error {
  args := m.Called(id, hide, updateDate)
  return args.Error(0)
}

func (m *ProviderMock) DeletePostById(id int) error {
  args := m.Called(id)
  return args.Error(0)
}

// ------------ USER ------------

func (m *ProviderMock) InsertUser(user entities.User) (*entities.User, error) {
  args := m.Called(user)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.User), args.Error(1)
}

func (m *ProviderMock) SelectUserByEmail(email string) (*entities.User, error) {
  args := m.Called(email)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.User), args.Error(1)
}

func (m *ProviderMock) SelectUserById(id int) (*entities.User, error) {
  args := m.Called(id)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.User), args.Error(1)
}

func (m *ProviderMock) SelectUserRulesByEmail(email string) (*int, error) {
  args := m.Called(email)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*int), args.Error(1)
}

func (m *ProviderMock) SelectUserRulesById(id int) (*int, error) {
  args := m.Called(id)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*int), args.Error(1)
}

func (m *ProviderMock) SelectUserPasswordByEmail(email string) (*string, error) {
  args := m.Called(email)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*string), args.Error(1)
}

func (m *ProviderMock) UpdateUserById(user entities.User) (*entities.User, error) {
  args := m.Called(user)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.User), args.Error(1)
}

func (m *ProviderMock) UpdateUserAdminRulesByEmail(email string, admin int) error {
  args := m.Called(email, admin)
  return args.Error(0)
}

func (m *ProviderMock) DeleteUserById(id int) error {
  args := m.Called(id)
  return args.Error(0)
}

func (m *ProviderMock) DeleteUserByEmail(email string) error {
  args := m.Called(email)
  return args.Error(0)
}

// ------------ COMMENTS ------------

func (m *ProviderMock) InsertComment(comment entities.Comment) (*entities.Comment, error) {
  args := m.Called(comment)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.Comment), args.Error(1)
}

func (m *ProviderMock) GetCommentsForPost(postId int, limit int, offset int) ([]*entities.Comment, error) {
  args := m.Called(postId)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).([]*entities.Comment), args.Error(1)
}


func (m *ProviderMock) UpdateComment(comment entities.Comment) (*entities.Comment, error) {
  args := m.Called(comment)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*entities.Comment), args.Error(1)
}

func (m *ProviderMock) DeleteComment(id int) error {
  args := m.Called(id)
  return args.Error(0)
}

func (m *ProviderMock) DeleteCommentsInPost(postId int) error {
  args := m.Called(postId)
  return args.Error(0)
}
