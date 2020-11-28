package osutil_test

import (
	"airdb.io/airdb/sailor/osutil"
	"testing"
)

func TestExec(t *testing.T) {
	bin := "echo"
	args := []string{"test"}

	osutil.Exec(bin, args)
}

func TestExecCommand(t *testing.T) {
	bin := "echo"
	args := []string{"test"}

	osutil.ExecCommand(bin, args)
}
