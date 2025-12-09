package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateComment(e echo.Context) error {
	var comment entities.Comment

	// Так как
	err := e.Bind(&comment)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// Валидол
	err = validator.New().Struct(comment)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Это на случай если с файлом всё нот окэ
	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	createdComment, err := s.uc.CreateComment(comment)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	// если пнг есть в комменте то мы его сохраняем
	if file != nil {
		err := s.saveCommentImage(file, createdComment.ID)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		}
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

// в связи с тем что файл в JSON не запихать мы дёргаем отдельную ручку
func (s *Server) GetCommentImage(e echo.Context) error {
	id := e.Param("id")

	path := fmt.Sprintf("%s/comment-%s.png", s.commentStaticPath, id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return e.JSON(http.StatusNotFound, "image not found")
	}

	return e.File(path)
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

// функция сохранения картинки на сервере
func (s *Server) saveCommentImage(file *multipart.FileHeader, id int) error {
	dst := fmt.Sprintf("%s/comment-%d.png", s.commentStaticPath, id)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//на случай если нету папки
	os.MkdirAll(s.commentStaticPath, 0755)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
