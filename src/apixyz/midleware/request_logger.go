package midleware

import (
	"apixyz/src/apixyz/util"
	"bytes"
	"io"
	"io/ioutil"
	"regexp"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

			re := regexp.MustCompile(`\r?\n`)
			var request = re.ReplaceAllString(readBody(rdr1), "")
			util.LogActivity(
				c.Request.Method,
				c.Request.URL,
				request,
				"HTTP Request Method",
				requestid.Get(c))
			c.Request.Body = rdr2
			c.Next()
		} else {
			util.LogActivity(
				c.Request.Method,
				c.Request.URL,
				"",
				"HTTP Request Method",
				requestid.Get(c))
			c.Next()
		}
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
