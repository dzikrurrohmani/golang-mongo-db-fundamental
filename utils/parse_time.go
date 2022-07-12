package utils

import "time"

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}