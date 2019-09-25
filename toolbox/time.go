package toolbox

import (
	"strconv"
	"time"
)

func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
