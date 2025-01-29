package api

import (
	"net/http"
	"strconv"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateComment(e echo.Context) error {
	var comment entities.Comment

	err := e.Bind(comment)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validator.New().Struct(comment)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	createdComment, err := s.uc.CreateComment(comment)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, createdComment)
}

func (s *Server) GetComments(e echo.Context) error {
	post_id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	commetns, err := s.uc.PostComments(post_id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, commetns)
}

func (s *Server) DeleteComment(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.uc.DeleteComment(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{
		"error": "ты не туда залез",
	})
}

func (s *Server) DeleteComments(e echo.Context) error {
	post_id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.uc.DeleteCommentsInPost(post_id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{
		"error": "ты не туда залез",
	})
}

func (s *Server) UpdateComment(e echo.Context) error {
	var comment entities.Comment

	err := e.Bind(comment)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validator.New().Struct(comment)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	comment.ID = id

	updatedComment, err := s.uc.UpdateComment(comment)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, updatedComment)
}
