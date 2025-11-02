package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
	"github.com/labstack/echo/v4"
)

// Обработчики связанные с постами
// Сделать что бы пост мог создаваться только админом
func (s *Server) CreatePost(e echo.Context) error {
	var post entities.Post

	err := e.Bind(&post)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	post.CreateDate = time.Now().Format("2006/01/02")
	post.UserId = UserIDFromToken(e)

	createdPost, err := s.uc.CreatePost(post)

	if err != nil {
		if errors.Is(err, entities.ErrPostNameConflict) {
			return e.JSON(http.StatusConflict, err.Error())
		}
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, createdPost)
}

func (s *Server) GetPosts(e echo.Context) error {
	posts, err := s.uc.ListPosts()

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, posts)
}

func (s *Server) GetPost(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	post, err := s.uc.SelectPost(id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, post)
}

// Сделать так что бы только админ поста мог обновлять пост
func (s *Server) UpdatePost(e echo.Context) error {
	var post entities.Post

	err := e.Bind(&post)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	post.ID = id

	post.UpdateDate = time.Now().Format("2006/01/02")

	updatedPost, err := s.uc.UpdatePost(post)

	if err != nil {
		if errors.Is(err, entities.ErrPostNameConflict) {
			return e.JSON(http.StatusConflict, err.Error())
		}
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, updatedPost)
}

// Сделать так что бы только админ поста мог удалять пост
func (s *Server) DeletePost(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.uc.DeletePost(id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{
		"error": "ты не туда залез",
	})
}

func (s *Server) MyPosts(e echo.Context) error {
	userId := UserIDFromToken(e)

	posts, err := s.uc.ListUserPosts(userId)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, posts)
}

func (s *Server) GetUserPosts(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	posts, err := s.uc.ListUserPosts(userId)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, posts)
}

func (s *Server) HidePost(e echo.Context) error {
	var post entities.Post

	err := e.Bind(&post)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	userId := UserIDFromToken(e)

	err = s.uc.HidePost(post.ID, userId)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{
		"error": "Пост скрыт с общей ленты публикаций",
	})
}
