package timeutil

import "time"

const (
	TimeFormatLong = "20060102150405"
	TimeFormatDay  = "2006-01-02"
)

const base = 1000

func UnixTimeMilliSecond() float64 {
	return float64(time.Now().UnixNano()/base) / base / base
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
