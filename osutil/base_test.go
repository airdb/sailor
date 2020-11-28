package osutil_test

import (
	"testing"

	"airdb.io/airdb/sailor/osutil"
	"github.com/stretchr/testify/assert"
)

func TestOS(t *testing.T) {
	if isw := osutil.IsWin(); isw {
		assert.True(t, isw)
		assert.False(t, osutil.IsMac())
		assert.False(t, osutil.IsLinux())
	}

	if ism := osutil.IsMac(); ism {
		assert.True(t, ism)
		assert.False(t, osutil.IsWin())
		assert.False(t, osutil.IsLinux())
	}

	if isl := osutil.IsLinux(); isl {
		assert.True(t, isl)
		assert.False(t, osutil.IsMac())
		assert.False(t, osutil.IsWin())
	}
}
