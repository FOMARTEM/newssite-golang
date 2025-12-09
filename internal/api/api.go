package api

import (
	"fmt"
	"os"
	"strings"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	server            *echo.Echo
	address           string
	commentStaticPath string

	secretKey string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase, secretKey string, frontAddress string, imagePath string) *Server {
	api := Server{
		uc:                uc,
		secretKey:         secretKey,
		commentStaticPath: imagePath,
	}

	api.server = echo.New()
	api.server.Logger.SetLevel(log.ERROR)

	logFile, err := os.OpenFile("../log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл log.txt: %v", err)
	}

	api.server.Logger.SetOutput(logFile)

	//формат логов
	api.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `[${time_custom}]  |  ${status}  |  ${method}  |  ${remote_ip}${path}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
		Output:           logFile,
	}))

	api.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontAddress},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	//  немного изменили что бы можно было попроще дёргать ручки без токена
	api.server.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {

			// Пути, которые игнорируют JWT
			allowed := map[string]bool{
				"/login":         true,
				"/signup":        true,
				"/posts":         true,
				"/comment/image": true,
			}

			// Полное совпадение
			if allowed[c.Path()] {
				return true
			}

			// Обработка динамического параметра /comment/image/:id
			if strings.HasPrefix(c.Path(), "/comment/image/") {
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
	api.server.PUT("/rules", api.EditRules)

	//посты
	api.server.POST("/post", api.CreatePost)
	api.server.GET("/posts", api.GetPosts)
	api.server.GET("/post/:id", api.GetPost)
	api.server.PUT("/post/:id", api.UpdatePost)
	api.server.DELETE("/post/:id", api.DeletePost)
	api.server.GET("/myposts", api.MyPosts)
	api.server.GET("/userposts/:id", api.GetUserPosts)
	api.server.PUT("/hidepost/:id", api.HidePost)

	//комментарии
	api.server.POST("/comment", api.CreateComment)
	api.server.GET("/comments/:id", api.GetComments)
	api.server.PUT("/comment/:id", api.UpdateComment)
	api.server.DELETE("/comment/:id", api.DeleteComment)
	api.server.DELETE("/comments/:id", api.DeleteComments)
	api.server.GET("/comment/image/:id", api.GetCommentImage)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}
