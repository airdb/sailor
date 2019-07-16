package sailor

import (
	"fmt"
	"os"
)

func ChmodAddUserPerm(filename string) (flag bool) {
	if err := os.Chmod(filename, 0744); err != nil {
		fmt.Println(err)
	}
	return
}

func MakeFileLink(filename, destfilename string) (flag bool) {
	err := os.Symlink(filename, destfilename)
	if err == nil {
		flag = true
	}
	return
}
