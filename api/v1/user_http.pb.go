package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/spf13/cast"
)

type UserServiceHTTPServer interface {
	Get(context.Context, *GetUserRequest) (*User, error)
	Create(context.Context, *CreateUserRequest) (*User, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) (*Empty, error)
}

type UserService struct {
	server UserServiceHTTPServer
	router gin.IRouter
	render render.Render
}

func RegisterUserServiceHTTPServer(srv UserServiceHTTPServer, router gin.IRouter, render render.Render) {
	s := UserService{server: srv, router: router, render: render}
	s.registerRouter()
}

func (s *UserService) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	user, err := s.server.Get(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, user)
}

func (s *UserService) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req.User); err != nil {
		s.render.Error(c, err)
		return
	}
	user, err := s.server.Create(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, user)
}

func (s *UserService) UpdateUser(c *gin.Context) {
	var (
		req UpdateUserRequest
		err error
	)
	if err = c.ShouldBindJSON(&req.User); err != nil {
		s.render.Error(c, err)
		return
	}
	if req.Id, err = cast.ToInt32E(c.Param("id")); err != nil {
		s.render.Error(c, err)
		return
	}
	user, err := s.server.Update(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, user)
}

func (s *UserService) DeleteUser(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	if _, err := s.server.Delete(c.Request.Context(), &req); err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, nil)
}

func (s *UserService) registerRouter() {
	s.router.GET("/user/:id", s.GetUser)
	s.router.POST("/user", s.CreateUser)
	s.router.PUT("/user/:id", s.UpdateUser)
	s.router.DELETE("/user/:id", s.DeleteUser)
}
