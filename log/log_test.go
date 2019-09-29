package log_test

import (
	"github.com/airdb/sailor/log"
	"testing"
)

func TestLogger(t *testing.T) {
	log := log.NewLogger(10000)
	log.SetLogger("console", "")
	log.Info("info")
	t.Log("hello")
}
