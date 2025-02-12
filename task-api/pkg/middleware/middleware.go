package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	internalerr "github.com/vygos/task/task-api/pkg/middleware/statuserr"
)

type HandlerErrorFunc func(g *gin.Context) error

func ErrorHandler(errorFunc HandlerErrorFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := errorFunc(c); err != nil {
			if errors.As(err, new(internalerr.ApiError)) {
				e := err.(internalerr.ApiError)
				c.AbortWithStatusJSON(e.StatusCode(), gin.H{
					"message": e.Error(),
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
	}
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
	if c.Request.Method == http.MethodOptions {
		c.Header("Access-Control-Allow-Headers", "content-type,*")
		c.AbortWithStatusJSON(http.StatusOK, nil)
	}
}
