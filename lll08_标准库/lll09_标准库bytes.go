package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := "hello world!"
	b1 := []byte("你好，世界！")
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("b1: %v\n", b1)

	//bytes和string的相互转换
	//1.bytes转string
	fmt.Printf("string(b1): %v\n", string(b1))
	//2.string转bytes
	fmt.Printf("[]byte(s1): %v\n", []byte(s1))

	//contains：检查bytes中是否包含子bytes
	fmt.Printf("bytes.Contains(b1, []byte(\"世界\")): %v\n", bytes.Contains(b1, []byte("世界")))

	//count：统计某个bytes出现的次数
	fmt.Printf("bytes.Count([]byte(s1), []byte(\"l\")): %v\n", bytes.Count([]byte(s1), []byte("l")))

	//compare:比较两个bytes The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	fmt.Printf("bytes.Compare(b1, []byte(\"hello\")): %v\n", bytes.Compare(b1, []byte("hello")))

	//分割bytes
	before, after, _ := bytes.Cut(b1, []byte("，"))
	fmt.Printf("before: %v\n", string(before))
	fmt.Printf("after: %v\n", string(after))

	//连接bytes
	b := bytes.Join([][]byte{b1, []byte(s1)}, []byte("==="))
	fmt.Printf("b: %v\n", string(b))

	//runes 转成utf8编码 这样能够正确计算中文长度
	r := bytes.Runes(b1)
	fmt.Printf("bytes.Runes(b1): %v\n", r)
	fmt.Printf("len(r): %v\n", len(r))

}
