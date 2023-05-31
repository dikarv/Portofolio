package util

import (
	"apixyz/src/apixyz/domain"
	"strconv"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	logger "github.com/sirupsen/logrus"
)

// HandleSuccess2 is a function of success process that send to the front as a response success
func HandleSuccess2(c *gin.Context, status int, code int, message string, data interface{}, detail interface{}, msg string) {
	// Assign struct ResponseWrapper with parameter above
	response := domain.Response2{
		Status:  strconv.Itoa(status),
		Message: message,
		Result:  data,
	}

	// Use JSON as a response and send initialize struct above
	c.JSON(code, response)

	reqID := requestid.Get(c)
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Info(msg)
	Logkoe.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
		"||detail =", detail)

}

// HandleSuccess is a function of success process that send to the front as a response success
func HandleSuccess(c *gin.Context, status int, code int, message string, data interface{}, detail interface{}, msg string) {
	// Assign struct ResponseWrapper with parameter above
	response := domain.Response{
		Status:  strconv.Itoa(status),
		Message: message,
		Data:    data,
	}

	// Use JSON as a response and send initialize struct above
	c.JSON(code, response)

	reqID := requestid.Get(c)
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Info(msg)
	Logkoe.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
		"||detail =", detail)

}

// HandleError is a function of failed process that send to the front as a response error
func HandleError(c *gin.Context, status int, code int, message string, detail interface{}, msg string) {
	// Assign struct ResponseWrapper with parameter above
	response := domain.Response{
		Status:  strconv.Itoa(status),
		Message: message,
		Data:    nil,
	}

	// Use JSON as a response and send initialize struct above
	c.JSON(code, response)

	reqID := requestid.Get(c)

	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Error(msg)
	Logkoe.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
		"||detail =", detail,
	)
}

func LogActivity(method interface{}, url interface{}, req interface{}, msg string, requestID string) {
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"method":  method,
		"url":     url,
		"request": req,
		"ID":      requestID,
	}).Info(msg)
	Logkoe.Info("Req  => method =", method,
		"||ID =", requestID,
		"||url =", url,
		"||param =", req)

}
