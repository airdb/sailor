package toolbox

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const (
	// ipRegular = "(\\d{1,3}\\.){3}\\d{1,3}/"
	ipRegular = "((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)/"
)

func GetLocalIP(devnet string) (ipaddr string) {
	cmd := exec.Command("/sbin/ip", "address", "show", devnet)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	reg := regexp.MustCompile(ipRegular)
	ipaddr = strings.TrimRight(reg.FindString(out.String()), "/")
	return
}
