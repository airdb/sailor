package sailor

import (
	"os"
	"runtime"
)

const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
	DevNull           = "/dev/null"
)

const GOOS string = runtime.GOOS
const GOARCH string = runtime.GOARCH

var GoVersion string = runtime.Version()
var OSPageSize int = os.Getpagesize()

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "NULL"
	}
	return hostname
}

var PID int = os.Getpid()
var PPID int = os.Getppid()

var UID int = os.Getuid()

func Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		dir = "./"
	}
	return dir
}
