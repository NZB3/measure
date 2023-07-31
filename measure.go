package measure

import (
	"fmt"
	"time"
)

func Measure(fn func()) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n\n", elapsed)
	return elapsed
}
