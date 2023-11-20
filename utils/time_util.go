package utils

import (
	"time"
)

const (
	apiTimeLayout = "2006-01-02T15:04:05"
)

func GetDateNow() string {
	now := time.Now().UTC()
	return now.Format(apiTimeLayout)
}
