package main

import (
	"fmt"
	"time"
)

func main() {

	// 每隔一秒打印一次当前时间
	t1 := time.NewTicker(time.Second)
	n := 0
	for _ = range t1.C {
		fmt.Printf("time.Now(): %v\n", time.Now())
		n += 1
		if n > 5 {
			t1.Stop()
			break
		}
	}

	t2 := time.NewTicker(time.Second)
	chanInt := make(chan int)
	go func() {
		//每秒钟向chanInt随机写一个数据
		for _ = range t2.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			case chanInt <- 4:
			case chanInt <- 5:

			}
		}

	}()
	sum := 0
	for v := range chanInt {
		sum += v
		fmt.Printf("v: %v\n", v)
		if sum >= 10 {
			close(chanInt)
			break
		}
	}

	print("main end...\n")

}
