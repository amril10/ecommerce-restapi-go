package errorhandler

import (
	"net/http"

	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/helper"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *ForbiddenError:
		statusCode = http.StatusForbidden
	}

	response := helper.Response(dto.ResponseWithParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
