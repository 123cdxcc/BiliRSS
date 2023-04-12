package pool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	pool := NewPool(10)
	for i := 0; i < 200; i++ {
		pool.Go(func(v ...interface{}) {
			fmt.Println("执行任务", v[0])
		}, i)
	}
	pool.Wait()
}
