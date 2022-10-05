package main

import (
	"fmt"
	"sync"
)

var n int = 100
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	lock.Lock() //加锁
	n += 1
	lock.Unlock() //解锁
	wg.Add(-1)
}

func sub() {
	lock.Lock()
	n -= 1
	lock.Unlock()
	wg.Add(-1) //完成协程任务，计数减一
}

func main() {

	//当多个协程共同处理某一个数据时，如果不进行加锁，就会出现数据混乱的情况，例如：多个协程对一个数值进行加一和减一的操作
	//因此在多协程中，我们可以对数据进行的操作加锁，让同一时刻只能有一个协程对数据进行操作
	for i := 0; i < 100000; i++ {
		wg.Add(1) //新增协程任务，计数加一
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait() //主协程等待子协程执行完毕
	fmt.Printf("n: %v\n", n)

}
