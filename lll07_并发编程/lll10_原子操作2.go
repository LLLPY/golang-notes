package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var i int32 = 100
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, 100) //加100
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -50) //减50
	fmt.Printf("i: %v\n", i)

	fmt.Printf("atomic.LoadInt32(&i): %v\n", atomic.LoadInt32(&i)) //读取i的值
	atomic.StoreInt32(&i, 1000)                                    //修改i的值为1000
	fmt.Printf("i: %v\n", i)

	//cas 其他操作的根基
	success := atomic.CompareAndSwapInt32(&i, i, 666) //修改i为666
	if success {

		fmt.Printf("修改成功！i: %v\n", i)
	}

	//swap 直接交换 比较暴力
	b := atomic.CompareAndSwapInt32(&i, i, 888)
	if b {
		fmt.Printf("修改成功！i: %v\n", i)
	}

}
