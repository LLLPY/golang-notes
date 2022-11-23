package main

import (
	"fmt"
	"os"
)

func main() {

	//1.获取所有环境变量
	fmt.Printf("os.Environ(): %v\n", os.Environ())
	// for _, v := range os.Environ() {
	// 	fmt.Printf("v: %v\n", v)
	// }

	//2.获取指定的环境变量
	fmt.Printf("os.Getenv(\"USER\"): %v\n", os.Getenv("USER"))

	//3.查找指定的环境变量
	s, b := os.LookupEnv("USER")
	if b {
		fmt.Printf("s: %v\n", s)
	}

	//4.设置环境变量
	os.Setenv("demo", "demo666")
	fmt.Printf("os.Getenv(\"demo\"): %v\n", os.Getenv("demo"))

	//5.替换环境变量 通过 $key 取到对应的环境变量的值
	fmt.Printf("os.ExpandEnv(\"user is ,demo is $demo.\"): %v\n", os.ExpandEnv("user is $USER,demo is ${demo}."))

	//6.清除所有环境变量(慎用)
	// os.Clearenv()
}
