package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.NewTimer(time.Second) //等待1秒钟
	fmt.Printf("time.Now(): %v\n", time.Now())
	res := <-t1.C //一直处于阻塞状态，直到到达等待时间
	fmt.Printf("res: %v\n", res)
	fmt.Printf("time.Now(): %v\n", time.Now())

	time.After(time.Second) //一秒后，After其实对NewTimer进行了一次封装
	fmt.Printf("time.Now(): %v\n", time.Now())

	//停止计时器
	t2 := time.NewTimer(time.Second * 2)

	go func() {
		print("执行到了这里...\n")
		<-t2.C
		print("匿名函数被执行...\n")
	}()

	t2.Stop() //如果调用了stop，上面的协程<-t2.C以后的语句都不会被执行
	time.Sleep(time.Second * 3)

	//修改计时时间
	t3 := time.NewTimer(time.Second)
	fmt.Printf("time.Now(): %v\n", time.Now())
	t3.Reset(time.Second * 2) //修改计时为2两秒
	<-t3.C
	fmt.Printf("time.Now(): %v\n", time.Now())

	fmt.Printf("main end...\n")
}
