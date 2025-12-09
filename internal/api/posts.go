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
// +
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

// +
func (s *Server) GetPosts(e echo.Context) error {
  limit, offset := s.getLimitOffset(e)
  posts, err := s.uc.ListPosts(limit, offset)

  if err != nil {
    return e.JSON(http.StatusInternalServerError, err.Error())
  }

  return e.JSON(http.StatusOK, posts)
}

// +
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

// +
func (s *Server) UpdatePost(e echo.Context) error {
  var post entities.Post
  var userId int

  err := e.Bind(&post)
  if err != nil {
    return e.JSON(http.StatusInternalServerError, err.Error())
  }

  id, err := strconv.Atoi(e.Param("id"))
  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  post.ID = id

  userId = UserIDFromToken(e)

  post.UserId = userId

  post.UpdateDate = time.Now().Format("2006/01/02")

  updatedPost, err := s.uc.UpdatePost(post, userId)

  if err != nil {
    if errors.Is(err, entities.ErrPostNameConflict) {
      return e.JSON(http.StatusConflict, err.Error())
    }
    return e.JSON(http.StatusInternalServerError, err.Error())
  }

  return e.JSON(http.StatusCreated, updatedPost)
}

// +
func (s *Server) DeletePost(e echo.Context) error {
  id, err := strconv.Atoi(e.Param("id"))
  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  userId := UserIDFromToken(e)

  err = s.uc.DeletePost(id, userId)

  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  return e.JSON(http.StatusOK, echo.Map{
    "error": "Пост успешно удалён",
  })
}

// +
func (s *Server) MyPosts(e echo.Context) error {
  userId := UserIDFromToken(e)

  limit, offset := s.getLimitOffset(e)

  posts, err := s.uc.ListUserPosts(userId, limit, offset)

  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  return e.JSON(http.StatusOK, posts)
}

// +
func (s *Server) GetUserPosts(e echo.Context) error {
  userId, err := strconv.Atoi(e.Param("id"))
  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  limit, offset := s.getLimitOffset(e)

  posts, err := s.uc.ListUserPosts(userId, limit, offset)

  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  return e.JSON(http.StatusOK, posts)
}

// +
func (s *Server) HidePost(e echo.Context) error {
  postId, err := strconv.Atoi(e.Param("id"))
  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  userId := UserIDFromToken(e)

  err = s.uc.HidePost(postId, userId)

  if err != nil {
    return e.JSON(http.StatusBadRequest, err.Error())
  }

  return e.JSON(http.StatusOK, echo.Map{
    "error": "Видимость поста изменена",
  })
}
