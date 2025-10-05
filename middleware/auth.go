package middleware

import (
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "Unauthorized, no token found"})
			c.Abort()
			return
		}

		claims, err := helper.VerifyToken(tokenString)
		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "Unauthorized, invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.ID)
		c.Set("roleID", claims.RoleID)
		c.Next()
	}
}
