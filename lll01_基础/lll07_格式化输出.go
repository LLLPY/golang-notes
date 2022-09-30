package main

import "fmt"

type studnet struct {
	Name string
}

func main() {

	//相应数据的默认格式
	a := 10
	fmt.Printf("a: %#v\n", a)

	tom := studnet{"Tom"}
	//数据的完整格式
	fmt.Printf("tom: %#v\n", tom)

	//%
	fmt.Printf("%%\n")

	//布尔占位符
	flag := true
	fmt.Printf("flag: %t\n", flag)

	//二进制输出
	fmt.Printf("%b\n", 10)

	//相应的Unicode码值对应的码符
	fmt.Printf("%c\n", 97)

	//八进制
	fmt.Printf("%o\n", 10)

	//十六进制 小写字母
	fmt.Printf("%x\n", 10)

	//十六进制 大写字母
	fmt.Printf("%X\n", 10)

	fmt.Printf("%q\n", 10)
	fmt.Printf("%U\n", 10)
	fmt.Printf("11111111111")
	//浮点数和复数
	fmt.Printf("%b\n", 10.1)
	fmt.Printf("%e\n", 10.1)
	fmt.Printf("%E\n", 10.1)
	fmt.Printf("2%5.2f2\n", 10.1)
	fmt.Printf("%g\n", 10.10000)
	fmt.Printf("%G\n", 10.000+2.0600i)

	//字符串
	fmt.Printf("[]byte{\"hello\", \"world\"}: %s\n", []byte("hello"))
	fmt.Printf("\"hello\": %q\n", "hello")
	fmt.Printf("\"hello\": %x\n", "hello")
	fmt.Printf("\"hello\": %X\n", "hello")

	//指针
	h := 100
	p_h := &h
	fmt.Printf("h: %p\n", &h)
	fmt.Printf("p_h: %p\n", p_h)

}
