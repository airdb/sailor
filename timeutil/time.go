package timeutil

import "time"

func UnixTimeMilliSecond() float64 {
	// nolint: gomnd
	return float64(time.Now().UnixNano()/1000) / 1000000
}

