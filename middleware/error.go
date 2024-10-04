package middleware

import (
	"net/http"

	customerror "github.com/SwaDeshiTech/kubesync/customError"
	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case customerror.Http:
				c.AbortWithStatusJSON(e.StatusCode, dto.Response{
					HttpStatus: e.StatusCode,
					Message:    e.Description,
					Response:   e.Response,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
					HttpStatus: http.StatusInternalServerError,
					Message:    "Service unavailable at this moment",
				})
			}
		}
	}
}
