package sailor

import (
	"testing"
	"time"
)

func TestProcessBar(t *testing.T) {
	var bar ProcessBar

	bar.NewOption(0, 100)

	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}

	bar.Finish()
}
