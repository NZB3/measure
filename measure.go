package measure

import (
	"time"
)

func Measure(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
