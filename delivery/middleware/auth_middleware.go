package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/utils/service"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	RequireToken(userRole ...string) gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(userRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aH authHeader
		if err := c.ShouldBindHeader(&aH); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		tokenString := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		claims, err := a.jwtService.VerifyAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("todo", claims)

		validRole := false
		for _, role := range userRole {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Resource"})
			return
		}
		c.Next()
	}
}

func (a *authMiddleware) RefreshToken() gin.HandlerFunc {
	panic("unimplemented")
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
