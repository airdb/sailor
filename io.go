package sailor

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadContentFromFile(filename string) (ret string) {
	fd, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		return
	}
	defer fd.Close()
	contentByte, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	ret = strings.Replace(string(contentByte), "\n", "", -1)
	return
}

func WriteStringToFile(wstring, filename string) (flag bool) {
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModeType)
	if err != nil {
		return
	}
	defer fd.Close()
	bytecount, err := fd.WriteString(wstring)
	if err != nil && 0 != bytecount {
		flag = true
	}
	return
}
