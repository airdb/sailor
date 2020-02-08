package sailor

import "time"

func UnixTimeMilliSecond() float64 {
	return float64(time.Now().UnixNano()/1000) / 1000000
}
