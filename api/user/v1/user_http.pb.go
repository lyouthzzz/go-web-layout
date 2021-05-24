package v1

import (
	context "context"

	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/go-web-layout/internal/render"
	"github.com/spf13/cast"
)

type UserHTTPService interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	Delete(context.Context, *DeleteUserRequest) (*Empty, error)
}

type UserHTTPServer struct {
	*xgin.Server
	service UserHTTPService
}

var ecodeRender = &render.EcodeRender{}

func NewUserHTTPServer(core *xgin.Server, service UserHTTPService) *UserHTTPServer {
	return &UserHTTPServer{Server: core, service: service}
}

func (srv *UserHTTPServer) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		return
	}
	user, err := srv.service.GetUser(c.Request.Context(), &req)
	if err != nil {
		ecodeRender.Error(c, err)
		return
	}

	ecodeRender.OK(c, user)
}

func (srv *UserHTTPServer) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ecodeRender.Error(c, err)
		return
	}
	user, err := srv.service.CreateUser(c.Request.Context(), &req)
	if err != nil {
		ecodeRender.Error(c, err)
		return
	}
	ecodeRender.OK(c, user)
}

func (s *UserHTTPServer) UpdateUser(c *gin.Context) {
	var (
		req UpdateUserRequest
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		ecodeRender.Error(c, err)
		return
	}
	if req.Id, err = cast.ToInt64E(c.Param("id")); err != nil {
		ecodeRender.Error(c, err)
		return
	}
	user, err := s.service.Update(c.Request.Context(), &req)
	if err != nil {
		ecodeRender.Error(c, err)
		return
	}
	ecodeRender.OK(c, user)
}

func (s *UserHTTPServer) DeleteUser(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ecodeRender.Error(c, err)
		return
	}
	if _, err := s.service.Delete(c.Request.Context(), &req); err != nil {
		ecodeRender.Error(c, err)
		return
	}
	ecodeRender.OK(c, nil)
}

func (srv *UserHTTPServer) BuildRoute() {
	srv.GET("/v1/user/:id", srv.GetUser)
	srv.POST("/v1/user", srv.CreateUser)
	srv.PUT("/v1/user", srv.UpdateUser)
	srv.DELETE("/v1/user", srv.DeleteUser)
}
