package osutil

import (
	"runtime"
)

// IsWin system. linux windows darwin.
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac system.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux system.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
