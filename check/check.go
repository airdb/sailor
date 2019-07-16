package check

import (
	"os"
	"reflect"
	"regexp"
	"runtime"
	"unsafe"
)

func IsEmptyValue(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func IsEmptyStruct(s struct{}) (flag bool) {
	if 0 != unsafe.Sizeof(s) {
		flag = true
	}
	return
}

func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// unix user id is 0 or not
func IsRoot() (flag bool) {
	if 0 == os.Getuid() {
		flag = true
	}
	return
}

func IsLinux() (flag bool) {
	if "linux" == runtime.GOOS {
		flag = true
	}
	return
}
func IsUnix() (flag bool) {
	if "unix" == runtime.GOOS {
		flag = true
	}
	return
}
func IsMac() (flag bool) {
	if "mac" == runtime.GOOS {
		flag = true
	}
	return
}
func IsWindows() (flag bool) {
	if "windows" == runtime.GOOS {
		flag = true
	}
	return
}

func IsInt(str string) bool {
	var reg = regexp.MustCompile(`^[-\\+]?[\\d]+$`)
	return reg.MatchString(str)
}

func IsChinese(str string) bool {
	var reg = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	return reg.MatchString(str)
}
