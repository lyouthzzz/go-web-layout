package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/spf13/cast"
)

type UserHTTPService interface {
	Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	Logout(context.Context, *UserLogoutRequest) (*Empty, error)
	Get(context.Context, *GetUserRequest) (*User, error)
	Create(context.Context, *CreateUserRequest) (*User, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) (*Empty, error)
}

type UserServer struct {
	service UserHTTPService
	render  render.Render
}

func NewUserHTTPServer(svc UserHTTPService, render render.Render) *UserServer {
	return &UserServer{service: svc, render: render}
}

func (s *UserServer) Login(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	loginResponse, err := s.service.Login(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	c.Header("session_id", loginResponse.Session.Id)
	s.render.OK(c, loginResponse.User)
}

func (s *UserServer) Logout(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		s.render.Error(c, err)
		return
	}
	if _, err := s.service.Logout(c.Request.Context(), &UserLogoutRequest{SessionId: sessionId}); err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, nil)
	return
}

func (s *UserServer) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	user, err := s.service.Get(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, user)
}

func (s *UserServer) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req.User); err != nil {
		s.render.Error(c, err)
		return
	}
	user, err := s.service.Create(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, user)
}

func (s *UserServer) UpdateUser(c *gin.Context) {
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
	user, err := s.service.Update(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, user)
}

func (s *UserServer) DeleteUser(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	if _, err := s.service.Delete(c.Request.Context(), &req); err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, nil)
}
