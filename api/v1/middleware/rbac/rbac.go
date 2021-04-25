package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/framework/pkg/rbac"
)

type Store interface {
	Roles(id string) ([]rbac.Role, error)
	Permissions(id string) ([]rbac.Permission, error)
}

func Middleware(rbacN *rbac.RBAC, store Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := ""
		roles, err := store.Roles(id)
		if err != nil {
			return
		}

		permission := rbac.NewUrlPermission("", c.Request.URL.Path, c.Request.Method)
		for _, r := range roles {
			_ = rbacN.AddRole(r)
			r.Permit(permission)
		}
	}
}
