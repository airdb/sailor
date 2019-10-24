package enum

import (
	"strings"
)

type Code uint

var AirdbSuccess uint = 20000
var AirdbFailed uint = 20001
var AirdbAuthFailed uint = 20002

var AirdbUnknown uint = 25000

var CodeMap = map[string]uint{
	"Success":      AirdbSuccess,
	"Failed":       AirdbFailed,
	"Auth failed":  AirdbAuthFailed,
	"Uknown error": AirdbUnknown,
}

var CodeMapInvert = InvertMap(CodeMap)

func FormCode(code uint) string {
	result, ok := CodeMapInvert[code]
	if ok {
		return result
	}
	return ""
}

func ToCode(sCode string) Code {
	result, ok := CodeMap[strings.ToLower(sCode)]
	if ok {
		return Code(result)
	}

	return 0
}

func InvertMap(input map[string]uint) map[uint]string {
	newMap := make(map[uint]string)
	for k, v := range input {
		newMap[v] = k
	}
	return newMap
}
