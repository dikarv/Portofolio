package midleware

import (
	"os"
	"time"

	logger "github.com/sirupsen/logrus"
)

func LoggingActivity() {
	dt := time.Now()
	date := dt.Format("20060102")
	var filename string = "apixyz/log/log" + date + ".log"

	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		logger.Fatal(err)
	}
	Formatter := new(logger.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	logger.SetFormatter(Formatter)
	logger.SetOutput(f)
}
