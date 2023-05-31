package auth

import (
	"apixyz/src/apixyz/database"
	"apixyz/src/apixyz/util"

	logger "github.com/sirupsen/logrus"
)

func CheckTokenID(tokenID string) bool {
	dbbafcommon, errdbCommon := database.GetConnectionSKS()
	if errdbCommon != nil {
		logger.WithFields(logger.Fields{
			"detail": errdbCommon,
		}).Error("Missing Connection")
		util.Logkoe.Info("msg =", "Missing Connection", "detail =", errdbCommon)
		return false
	}
	rows, errquery := dbbafcommon.Query("SELECT * FROM TOKEN_USR WHERE TOKEN_ID = ?", tokenID)

	defer rows.Close()
	if errquery != nil {
		logger.WithFields(logger.Fields{
			"detail": errquery,
		}).Error("Missing Query Statement")
		util.Logkoe.Info("msg =", "Missing Connection", "detail =", errquery)
		return false
	}
	if rows.Next() {
		logger.WithFields(logger.Fields{
			"detail": "Token ID found",
		}).Info("Token ID found")
		util.Logkoe.Info("msg =", "Token ID found")
		return true
	}

	logger.WithFields(logger.Fields{
		"detail": "Token ID not found",
	}).Info("Token ID not found")
	util.Logkoe.Info("msg =", "Token ID not found")
	return false
}
