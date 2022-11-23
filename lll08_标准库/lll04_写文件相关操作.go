package main

import (
	"fmt"
	"os"
)

// 以字节的形式写入
func write_as_byte() {
	f, err := os.OpenFile("a.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.Write([]byte("hello golang!"))
	f.Close()
}

//以字符串的形式写入
func write_as_string() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0777) //以追加的形式写入
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteString("hello world!")
	f.Close()
}

//随机位置写入
func write_random_position() {
	f, err := os.OpenFile("a.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteAt([]byte("你好，中国！"), 1)
	f.Close()
}

func main() {
	write_as_byte()
	write_as_string()
	write_random_position()
}
