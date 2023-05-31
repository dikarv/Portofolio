package midleware

import (
	"apixyz/src/apixyz/util"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			err := recover()

			// check if error panic is exist
			if err != nil {
				if ne, ok := err.(error); ok {
					errWrap := errors.WithStack(ne)
					stackStr := fmt.Sprintf("%+v", errWrap)
					//Write log
					logger.WithFields(logger.Fields{
						"URL":    c.Request.URL,
						"Method": c.Request.Method,
						"Error":  stackStr,
					}).Error("Failed make request")
					util.Logkoe.Info("Req to ", c.Request.URL, " | Method=", c.Request.Method, " | Error =>", stackStr)
				} else {
					//Write log
					logger.WithFields(logger.Fields{
						"URL":    c.Request.URL,
						"Method": c.Request.Method,
						"Error":  err,
					}).Error("Failed make request")
					util.Logkoe.Info("Req to ", c.Request.URL, " | Method=", c.Request.Method, " | Error =>", err)
				}
				//Write message and abort the request
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Mohon Maaf, Terjadi kesalahan pada sistem, silahkan coba beberapa saat lagi",
					"status":  http.StatusInternalServerError,
					"data":    nil,
				})
				c.Abort()
			}
		}()

		// execute request
		c.Next()
	}
}
