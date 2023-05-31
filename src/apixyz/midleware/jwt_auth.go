package midleware

import (
	authRepo "apixyz/src/apixyz/auth"
	"apixyz/src/apixyz/domain"
	"apixyz/src/apixyz/util"
	"apixyz/src/config"
	"log"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func JWTMiddleware() (gin.HandlerFunc, *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            config.Realm(),
		Key:              []byte(config.SuperSecretPassword()),
		Timeout:          5 * time.Second,
		MaxRefresh:       10 * time.Hour,
		SigningAlgorithm: "HS256",
		IdentityKey:      "tokenid",
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &domain.TokenClaims{
				TokenID:      claims["tokenid"].(string),
				UserID:       claims["userid"].(string),
				IdentityUser: claims["identityuser"].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			isTokenIDExist := authRepo.CheckTokenID(claims["tokenid"].(string))
			if isTokenIDExist {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"response": code,
				"result":   nil,
				"message":  strings.Title(message),
			})
			logger.WithFields(logger.Fields{
				"detail": "Unauthorized",
			}).Info("Unauthorized")
			util.Logkoe.Info("msg =", message)
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	// }, authMiddleware
	return authMiddleware.MiddlewareFunc(), authMiddleware
}
