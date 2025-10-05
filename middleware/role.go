package middleware

import (
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exist := c.Get("roleID")
		if !exist {
			errorhandler.HandleError(c, &errorhandler.ForbiddenError{Message: "Forbidden, no role found"})
			c.Abort()
			return
		}

		userRole := roleID.(int)

		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		errorhandler.HandleError(c, &errorhandler.ForbiddenError{Message: "Forbidden, insufficient permission"})
		c.Abort()
	}
}
