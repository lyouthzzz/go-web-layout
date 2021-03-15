package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/spf13/cast"
)

type UserServiceHTTPServer interface {
	Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	Logout(context.Context, *UserLogoutRequest) (*Empty, error)
	Get(context.Context, *GetUserRequest) (*User, error)
	Create(context.Context, *CreateUserRequest) (*User, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) (*Empty, error)
}

type UserService struct {
	server UserServiceHTTPServer
	router *gin.RouterGroup
	render render.Render
}

func RegisterUserServiceHTTPServer(srv UserServiceHTTPServer, routerGroup *gin.RouterGroup, render render.Render) {
	s := UserService{server: srv, router: routerGroup, render: render}
	s.registerRouter()
}

func (s *UserService) Login(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	loginResponse, err := s.server.Login(c.Request.Context(), &req)
	if err != nil {
		s.render.Error(c, err)
		return
	}
	c.Header("session_id", loginResponse.Session.Id)
	s.render.OK(c, loginResponse.User)
}

func (s *UserService) Logout(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		s.render.Error(c, err)
		return
	}
	if _, err := s.server.Logout(c.Request.Context(), &UserLogoutRequest{SessionId: sessionId}); err != nil {
		s.render.Error(c, err)
		return
	}
	s.render.OK(c, nil)
	return
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
	s.router.POST("/user/login", s.Login)
	s.router.POST("/user/logout", s.Logout)
	s.router.GET("/user/:id", s.GetUser)
	s.router.POST("/user", s.CreateUser)
	s.router.PUT("/user/:id", s.UpdateUser)
	s.router.DELETE("/user/:id", s.DeleteUser)
}
