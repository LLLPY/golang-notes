package main

import (
	"fmt"
	"runtime"
)

func show_msg(msg int) {
	for i := 0; i < 50; i++ {
		if i == 3 { //i等于3时直接结束协程
			defer fmt.Print("hello world!") //defered的语句在退出前会被执行

			runtime.Goexit()
		}
		fmt.Printf("msg: %v-%v\n", msg, i)
		// time.Sleep(time.Millisecond * 100)

	}

}

func main() {
	runtime.GOMAXPROCS(1) //设置核心数
	go show_msg(100)
	go show_msg(200)
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.Gosched()
	runtime.Gosched()
	// time.Sleep(time.Millisecond * 1000)
	print("main end...")
}
