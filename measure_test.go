package measure

import (
	"fmt"
	"testing"
	"time"
)

func TestMeasure(t *testing.T) {
	Measure(func() {
		for i := 0; i < 100000000; i++ {
			_ = i
		}
	})
	fmt.Println("TestMeasure done")
	time.Sleep(10 * time.Second)
}
