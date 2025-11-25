package usecase_test

import (
	"errors"
	"testing"

	"github.com/FOMARTEM/newssite-golang/internal/provider/mocks"
	"github.com/FOMARTEM/newssite-golang/internal/usecase"
	"github.com/stretchr/testify/assert"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

func TestUsecase_CreateComment(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	comment := entities.Comment{
		ID:          1,
		CommentText: "Nice post",
		PostId:      10,
		UserId:      5,
	}

	t.Run("error while checking user rules", func(t *testing.T) {
		mockP.On("SelectUserRulesById", 5).Return(nil, errors.New("db error"))

		res, err := u.CreateComment(comment)

		assert.Nil(t, res)
		assert.EqualError(t, err, "db error")
	})

	t.Run("user rules too low", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		rules := 0

		mockP.On("SelectUserRulesById", 5).Return(&rules, nil)

		res, err := u.CreateComment(comment)

		assert.Nil(t, res)
		assert.Equal(t, entities.ErrCommentNotFound, err)
	})

	t.Run("success", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		rules := 2

		mockP.On("SelectUserRulesById", 5).Return(&rules, nil)
		mockP.On("InsertComment", comment).Return(&comment, nil)

		res, err := u.CreateComment(comment)

		assert.NoError(t, err)
		assert.Equal(t, &comment, res)
	})

	t.Run("insert error", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		rules := 3

		mockP.On("SelectUserRulesById", 5).Return(&rules, nil)
		mockP.On("InsertComment", comment).Return(nil, errors.New("insert failed"))

		res, err := u.CreateComment(comment)

		assert.Nil(t, res)
		assert.EqualError(t, err, "insert failed")
	})
}

func TestUsecase_PostComments(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	comments := []*entities.Comment{
		{ID: 1, CommentText: "Hello"},
	}

	t.Run("success", func(t *testing.T) {
		mockP.On("GetCommentsForPost", 10).Return(comments, nil)

		res, err := u.PostComments(10)

		assert.NoError(t, err)
		assert.Equal(t, comments, res)
	})

	t.Run("db error", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		mockP.On("GetCommentsForPost", 20).Return(nil, errors.New("db error"))

		res, err := u.PostComments(20)

		assert.Nil(t, res)
		assert.EqualError(t, err, "db error")
	})
}
func TestUsecase_UpdateComment(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	comment := entities.Comment{ID: 1, CommentText: "Edit"}

	t.Run("db error", func(t *testing.T) {
		mockP.On("UpdateComment", comment).Return(nil, errors.New("update error"))

		res, err := u.UpdateComment(comment)

		assert.Nil(t, res)
		assert.EqualError(t, err, "update error")
	})

	t.Run("comment not found", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		mockP.On("UpdateComment", comment).Return(nil, nil)

		res, err := u.UpdateComment(comment)

		assert.Nil(t, res)
		assert.Equal(t, entities.ErrCommentNotFound, err)
	})

	t.Run("success", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		updated := &entities.Comment{ID: 1, CommentText: "Updated"}

		mockP.On("UpdateComment", comment).Return(updated, nil)

		res, err := u.UpdateComment(comment)

		assert.NoError(t, err)
		assert.Equal(t, updated, res)
	})
}
func TestUsecase_DeleteComment(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	t.Run("success", func(t *testing.T) {
		mockP.On("DeleteComment", 1).Return(nil)

		err := u.DeleteComment(1)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		mockP.On("DeleteComment", 2).Return(errors.New("db error"))

		err := u.DeleteComment(2)

		assert.EqualError(t, err, "db error")
	})
}
func TestUsecase_DeleteCommentsInPost(t *testing.T) {
	mockP := new(mocks.ProviderMock)
	u := usecase.NewUsecase(mockP)

	t.Run("success", func(t *testing.T) {
		mockP.On("DeleteCommentsInPost", 10).Return(nil)

		err := u.DeleteCommentsInPost(10)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockP.ExpectedCalls = nil
		mockP.On("DeleteCommentsInPost", 20).Return(errors.New("err delete"))

		err := u.DeleteCommentsInPost(20)

		assert.EqualError(t, err, "err delete")
	})
}
