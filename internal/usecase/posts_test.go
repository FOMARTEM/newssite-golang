package usecase_test

import (
	"testing"

	"github.com/FOMARTEM/newssite-golang/internal/provider/mocks"
	"github.com/FOMARTEM/newssite-golang/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

func TestCreatePost(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	post := entities.Post{UserId: 1, Name: "Test", Text: "Lorem ipsum dolor sit amet...."}

	t.Run("access denied", func(t *testing.T) {
		rules := 3
		mockP.On("SelectUserRulesById", post.UserId).Return(&rules, nil)

		res, err := u.CreatePost(post)
		assert.Nil(t, res)
		assert.Equal(t, entities.ErrAccessDenied, err)
	})

	t.Run("success", func(t *testing.T) {
		rules := 7
		mockP.ExpectedCalls = nil
		mockP.On("SelectUserRulesById", post.UserId).Return(&rules, nil)
		mockP.On("InsertPost", post).Return(&post, nil)

		res, err := u.CreatePost(post)
		assert.NoError(t, err)
		assert.Equal(t, &post, res)
	})
}

func TestSelectPost(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	post := &entities.Post{ID: 10}

	t.Run("not found", func(t *testing.T) {
		mockP.On("SelectPostById", 10).Return(nil, nil)

		res, err := u.SelectPost(10)
		assert.Nil(t, res)
		assert.Equal(t, entities.ErrPostNotFound, err)
	})

	t.Run("success", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		mockP.On("SelectPostById", 10).Return(post, nil)

		res, err := u.SelectPost(10)
		assert.NoError(t, err)
		assert.Equal(t, post, res)
	})
}

func TestListPosts(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	posts := []*entities.Post{{ID: 1}}

	mockP.On("SelectAllPosts").Return(posts, nil)

	res, err := u.ListPosts()
	assert.NoError(t, err)
	assert.Equal(t, posts, res)
}

func TestListUserPosts(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	posts := []*entities.Post{{ID: 1}}

	mockP.On("SelectUserPosts", 5).Return(posts, nil)

	res, err := u.ListUserPosts(5)
	assert.NoError(t, err)
	assert.Equal(t, posts, res)
}

func TestUpdatePost(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	post := entities.Post{ID: 1}
	dbPost := &entities.Post{ID: 1, UserId: 7}

	t.Run("wrong owner", func(t *testing.T) {
		mockP.On("SelectPostById", 1).Return(dbPost, nil)

		res, err := u.UpdatePost(post, 999)
		assert.Nil(t, res)
		assert.Equal(t, entities.ErrPostNotFound, err)
	})

	t.Run("access denied", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		rules := 2

		mockP.On("SelectPostById", 1).Return(dbPost, nil)
		mockP.On("SelectUserRulesById", 7).Return(&rules, nil)

		res, err := u.UpdatePost(post, 7)
		assert.Nil(t, res)
		assert.Equal(t, entities.ErrAccessDenied, err)
	})

	t.Run("success", func(t *testing.T) {
		mockP.ExpectedCalls = nil

		rules := 7
		updated := &entities.Post{ID: 1, Name: "AAA"}

		mockP.On("SelectPostById", 1).Return(dbPost, nil)
		mockP.On("SelectUserRulesById", 7).Return(&rules, nil)
		mockP.On("UpdatePostById", post).Return(updated, nil)

		res, err := u.UpdatePost(post, 7)
		assert.NoError(t, err)
		assert.Equal(t, updated, res)
	})
}

func TestHidePost(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	post := &entities.Post{ID: 1, Hide: 0}

	t.Run("access denied", func(t *testing.T) {
		rules := 3
		mockP.On("SelectUserRulesById", 5).Return(&rules, nil)

		err := u.HidePost(1, 5)
		assert.Equal(t, entities.ErrAccessDenied, err)
	})

	t.Run("success hide/unhide", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		rules := 7

		mockP.On("SelectUserRulesById", 5).Return(&rules, nil)
		mockP.On("SelectPostById", 1).Return(post, nil)
		mockP.On("EditVisibilityById", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		err := u.HidePost(1, 5)
		assert.NoError(t, err)
	})
}

func TestDeletePost(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	dbPost := &entities.Post{ID: 10, UserId: 7}

	t.Run("wrong owner", func(t *testing.T) {
		mockP.On("SelectPostById", 10).Return(dbPost, nil)

		err := u.DeletePost(10, 999)
		assert.Equal(t, entities.ErrPostNotFound, err)
	})

	t.Run("access denied", func(t *testing.T) {
		mockP.ExpectedCalls = nil

		rules := 1

		mockP.On("SelectPostById", 10).Return(dbPost, nil)
		mockP.On("SelectUserRulesById", 7).Return(&rules, nil)

		err := u.DeletePost(10, 7)
		assert.Equal(t, entities.ErrAccessDenied, err)
	})

	t.Run("success", func(t *testing.T) {
		mockP.ExpectedCalls = nil

		rules := 7

		mockP.On("SelectPostById", 10).Return(dbPost, nil)
		mockP.On("SelectUserRulesById", 7).Return(&rules, nil)
		mockP.On("DeleteCommentsInPost", 10).Return(nil)
		mockP.On("DeletePostById", 10).Return(nil)

		err := u.DeletePost(10, 7)
		assert.NoError(t, err)
	})
}
