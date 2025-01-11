package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateUser(e echo.Context) error {
	var user entities.User

	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validator.New().Struct(user)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	createdUser, err := s.uc.CreateUser(user)

	if err != nil {
		if errors.Is(err, entities.ErrUserNameConflict) ||
			errors.Is(err, entities.ErrUserEmailConflict) ||
			errors.Is(err, entities.ErrUserAlreadyExist) {
			return e.JSON(http.StatusConflict, err.Error())
		}
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	createdUser.Password = ""

	return e.JSON(http.StatusCreated, createdUser)
}

func (s *Server) Login(e echo.Context) error {
	var user entities.User

	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	login, err := s.uc.CheckPasswordUser(user)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if !*login {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "ну ты постарался чё сказать, сломал не пойми что",
		})
	}

	u, err := s.uc.SelectUserByEmail(user.Email)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	u.Token, err = token.SignedString([]byte(s.secretKey))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	u.Password = ""

	return e.JSON(http.StatusOK, u)
}

func (s *Server) GetUser(e echo.Context) error {
	userId := UserIDFromToken(e)

	user, err := s.uc.SelectUserByID(userId)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = ""

	return e.JSON(http.StatusOK, user)
}

// сделать так что бы проверялось совпадение старого пароля, а уже потом только измененеие данных пользователя
func (s *Server) UpdateUser(e echo.Context) error {
	var user entities.User

	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validator.New().Struct(user)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	updateUser, err := s.uc.UpdateUser(user)

	if err != nil {
		if errors.Is(err, entities.ErrUserNameConflict) ||
			errors.Is(err, entities.ErrUserEmailConflict) ||
			errors.Is(err, entities.ErrUserAlreadyExist) {
			return e.JSON(http.StatusConflict, err.Error())
		}
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	updateUser.Password = ""

	return e.JSON(http.StatusOK, updateUser)
}

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

func UserIDFromToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return int(claims["id"].(float64))
}
