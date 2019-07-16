package log_test

import (
	"github.com/airdb/sailor/log"
	"testing"
)

func TestLogger(t *testing.T) {
	log.Logger().Println("hello")
	t.Log("hello")
}
