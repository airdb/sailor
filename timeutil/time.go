package timeutil

import "time"

const (
	TimeFormatLong = "20060102150405"
	TimeFormatDay  = "2006-01-02"
)

func UnixTimeMilliSecond() float64 {
	// nolint: gomnd
	return float64(time.Now().UnixNano()/1000) / 1000000
}

func GetNowLong() string {
	return time.Now().Format(TimeFormatLong)
}

func GetYearMonthDay() string {
	return time.Now().Format(TimeFormatDay)
}

func GetTimeRFC(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(time.RFC3339)
}
