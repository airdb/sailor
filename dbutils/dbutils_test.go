package dbutils

import (
	"testing"
)

func TestInitDefault(t *testing.T) {
	// os.Setenv("GDBC", "root:hello@tcp(127.0.0.1:3306)/dev_mina")
	InitDefault()
}
