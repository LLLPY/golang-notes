package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var N int32 = 100
var wg sync.WaitGroup

func add() {
	atomic.AddInt32(&N, 1) //加一操作
	defer wg.Add(-1)
}

func sub() {
	atomic.AddInt32(&N, -1) //减一操作
	defer wg.Add(-1)
}
func main() {
	for i := 0; i < 100000; i++ {
		go add()
		wg.Add(1)
		go sub()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("N: %v\n", N)
}
