package enum

var AirdbSuccess uint = 20000
var AirdbFailed uint = 20001
var AirdbAuthFailed uint = 20002

var AirdbUnknown uint = 24999
var AirdbUndefined uint = 25000

var CodeMap = map[uint]string{
	AirdbSuccess:    "Success",
	AirdbFailed:     "Failed",
	AirdbAuthFailed: "Auth Failed",
	AirdbUnknown:    "Uknown Error",
	AirdbUndefined:  "Undefined Error",
}

var CodeMapInvert = InvertMap(CodeMap)

func FormCode(code uint) string {
	result, ok := CodeMap[code]
	if ok {
		return result
	}
	return CodeMap[AirdbUndefined]
}

func ToCode(sCode string) uint {
	result, ok := CodeMapInvert[sCode]
	if ok {
		return result
	}

	return AirdbUndefined
}

func InvertMap(input map[uint]string) map[string]uint {
	newMap := make(map[string]uint)
	for k, v := range input {
		newMap[v] = k
	}
	return newMap
}
