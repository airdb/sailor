package timeutil_test

import (
	"testing"
	"time"

	"github.com/airdb/sailor/timeutil"
)

func TestGetNowLong(t *testing.T) {
	timeutil.GetNowLong()
}

func TestGetYearMonthDay(t *testing.T) {
	timeutil.GetYearMonthDay()
}

func TestGetTimeRFC(t *testing.T) {
	now := time.Now().Unix()
	timeutil.GetTimeRFC(now)
}
