package dbutils

import (
	"testing"
)

func TestInitDefault(t *testing.T) {
	InitDefault()
	InitDB("")
	DefaultDB().Get("dev_mina")
	DefaultDB().Get("dev_mina")
	DefaultDB()
}
