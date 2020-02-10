package enum

import (
	"strings"
)

// Airdb error codes for user.
const (
	AirdbSuccess    uint = 20000
	AirdbFailed     uint = 20001
	AirdbAuthFailed uint = 20002
	AirdbUndefined  uint = 24999
	AirdbUnknown    uint = 25000
)

var CodeMap = map[uint]string{
	AirdbSuccess:    "Success",
	AirdbFailed:     "Failed",
	AirdbAuthFailed: "Auth failed",
	AirdbUndefined:  "Undefined",
	AirdbUnknown:    "Uknown error",
}

func FormCode(code uint) string {
	if result, ok := CodeMap[code]; ok {
		return result
	}

	return CodeMap[AirdbUnknown]
}

func ToCode(sCode string) uint {
	for k, v := range CodeMap {
		if v == strings.ToLower(sCode) {
			return k
		}
	}

	return AirdbUnknown
}
