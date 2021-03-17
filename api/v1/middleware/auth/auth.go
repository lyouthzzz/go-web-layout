package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/framework/pkg/auth/authn"
)

func Middleware(sessionAuthN authn.Authenticator) gin.HandlerFunc {

	return authn.NewAuthenticator(sessionAuthN)
}
