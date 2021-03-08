package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserServiceHTTPServer interface {
	Get(context.Context, *GetUserRequest) (*User, error)
	Create(context.Context, *CreateUserRequest) (*User, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) (*empty.Empty, error)
}

type Render interface {
	Error(c *gin.Context, err error)
	OK(c *gin.Context, data interface{})
}

type UserService struct {
	server UserServiceHTTPServer
	router gin.IRouter
	render Render
}

func RegisterUserServiceHTTPServer(srv UserServiceHTTPServer, router gin.IRouter) {
	s := UserService{server: srv, router: router}
	s.registerRouter()
}

func (s *UserService) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		s.render.Error(c, err)
		return
	}
	// 业务逻辑

	// 成功响应
	s.render.OK(c, &User{Name: "成功"})
}

func (s *UserService) CreateUser(c *gin.Context) {
}

func (s *UserService) UpdateUser(c *gin.Context) {
}

func (s *UserService) DeleteUser(c *gin.Context) {

}

func (s *UserService) registerRouter() {
	s.router.GET("/user/:id", s.GetUser)
	s.router.POST("/user", s.CreateUser)
	s.router.PUT("/user/:id", s.UpdateUser)
	s.router.DELETE("/user/:id", s.DeleteUser)
}
