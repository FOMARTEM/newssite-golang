package api

import (
	"fmt"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	server  *echo.Echo
	address string

	secretKey string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase, secretKey string) *Server {
	api := Server{
		uc:        uc,
		secretKey: secretKey,
	}

	api.server = echo.New()
	api.server.Logger.SetLevel(log.ERROR)

	api.server.Use(middleware.Logger())

	api.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},                                                                // Разрешённые источники (React клиент)
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},                                             // Разрешённые HTTP методы
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // Разрешённые заголовки
	}))

	api.server.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" || c.Path() == "/posts" {
				return true
			}
			return false
		},
	}))

	//пользователь
	api.server.POST("/signup", api.CreateUser)
	api.server.POST("/login", api.Login)
	api.server.GET("/profile", api.GetUser)
	api.server.PUT("/profile", api.UpdateUser)
	//может быть когда нибудь
	//api.server.PUT("/rules", api.Editrules)

	//посты
	api.server.POST("/post", api.CreatePost)
	api.server.GET("/posts", api.GetPosts)
	api.server.GET("/post/:id", api.GetPost)
	api.server.PUT("/post/:id", api.UpdatePost)
	api.server.DELETE("/post/:id", api.DeletePost)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}
