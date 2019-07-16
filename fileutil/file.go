package fileutil

import (
	"strings"
)

// return the source filename after the last slash
func ChopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
