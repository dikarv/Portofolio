package midleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// timeout middleware wraps the request context with a timeout
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded && !c.Writer.Written() {

				// write response and abort the request
				c.JSON(http.StatusGatewayTimeout, gin.H{
					"status":  http.StatusGatewayTimeout,
					"message": "Server Timeout",
				})
				c.Abort()
			}

			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
