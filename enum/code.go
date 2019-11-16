package enum

import (
	"strings"
)

type Code uint

const (
	AirdbSuccess    Code = 20000
	AirdbFailed     Code = 20001
	AirdbAuthFailed Code = 20002
	AirdbUnknown    Code = 25000
)

var CodeMap = map[Code]string{
	AirdbSuccess:    "Success",
	AirdbFailed:     "Failed",
	AirdbAuthFailed: "Auth failed",
	AirdbUnknown:    "Uknown error",
}

func FormCode(code Code) string {
	if result, ok := CodeMap[code]; ok {
		return result
	}
	return CodeMap[AirdbUnknown]
}

func ToCode(sCode string) Code {
	for k, v := range CodeMap {
		if v == strings.ToLower(sCode) {
			return k
		}
  }
	return AirdbUnknown
}
