package helper

import (
	"time"
)

func GetCurrentTime() string {
	currentDate := time.Now()
	currentDataFormatted := currentDate.Format("2006-01-02")
	return currentDataFormatted
}
