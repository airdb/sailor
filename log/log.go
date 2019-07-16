package log

import (
	"log"
	"os"
)

func init() {
}

func Logger() *log.Logger {
	logger := log.New(os.Stdin, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
