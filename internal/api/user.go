package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// +
func (s *Server) CreateUser(e echo.Context) error {
	var user entities.User

	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
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

// +
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
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()

	u.Token, err = token.SignedString([]byte(s.secretKey))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	u.Password = ""

	return e.JSON(http.StatusOK, u)
}

// +
func (s *Server) GetUser(e echo.Context) error {
	userId := UserIDFromToken(e)

	user, err := s.uc.SelectUserByID(userId)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = ""

	return e.JSON(http.StatusOK, user)
}

// +
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

	user.ID = UserIDFromToken(e)

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

// +
func (s *Server) EditRules(e echo.Context) error {
	adminId := UserIDFromToken(e)
	var user entities.User

	adminUser, err := s.uc.SelectUserByID(adminId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	adminRules := adminUser.AdminRole

	if adminRules != 7 {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "У вас нету доступна на обновление",
		})
	}

	err = e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	flag, err := s.uc.UpdateAdminRules(user.Email, user.AdminRole)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	} else if !*flag {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Непредвиденная ошибка",
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message": "Обновление прав прошло успешно",
	})

}

func UserIDFromToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	//fmt.Print(user)
	claims := user.Claims.(jwt.MapClaims)
	//fmt.Print(claims)
	return int(claims["id"].(float64))
}
