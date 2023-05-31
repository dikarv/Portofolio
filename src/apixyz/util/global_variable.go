package util

import (
	"github.com/alexcesaro/log/stdlog"
	"time"
)

var Logkoe = stdlog.GetFromFlags()

const (
	AGENT_STATUS_UNACTIVE = 0
	AGENT_STATUS_ACTIVE   = 1
	AGENT_STATUS_SUSPEND  = 2
	NO_FLAG               = 99

	AGENT_NOT_YET_APPROVE = 0
	AGENT_APPROVE         = 1
	AGENT_REJECT          = 2
	AGENT_REVISION        = 3

	ERR_NOT_FOUND         = "Data Not Found"
	ERR_CONFINS_NOT_FOUND = "Data Tidak Ada Pada Sistem"
	ERR_BAD_REQUEST       = "Bad Request"
	ERR_GENERAL           = "Maaf terjadi kesalahan pada sistem, silahkan coba lagi"

	RELIGION_MASTER_TYPE_CODE  = "RELIGION"
	EDUCATION_MASTER_TYPE_CODE = "EDUCATION"
	MARITAL_MASTER_TYPE_CODE   = "MARITAL_STAT"

	DATE_ONLY_LAYOUT = "2006-01-02"

	SQL_DATE_LAYOUT = "2006-01-02T15:04:05Z"

	DATETIME_ONLY_LAYOUT = "2006-01-02 15:04:05"
)

var Zone, _ = time.LoadLocation("Asia/Jakarta")
