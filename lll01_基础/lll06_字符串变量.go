package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//单行字符串 支持转义
	a := "hello world"

	//多行字符串 不支持转义
	b := `
			<div>
			<p>hello</p>
			</div>
		`
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	//字符串的拼接
	s1 := "hello"
	s2 := "world"
	//1.加号拼接
	res1 := s1 + " " + s2
	fmt.Printf("s: %v\n", res1)

	//2.字符串格式化 Sprintf
	res2 := fmt.Sprintf("%s %s", s1, s2)
	fmt.Printf("res2: %v\n", res2)

	//3.strings.join()
	/*
		join会根据字符串数组得到内容，计算出一个拼接之后的字符串的长度，然后申请对应大小的内存，一个一个字符填入，
		在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数组的代价也很高
	*/
	res3 := strings.Join([]string{s1, s2}, " ") //第一个参数是字符串数组，第二个参数是连接符
	fmt.Printf("res3: %v\n", res3)

	//4.buffer.WriteString()
	/*
		这个比较理想，可以当成可变字符串使用，对内存的增长也有优化，如果能预估字符串的长度，还可以使用buffer.Grow()接口来设置capacity
	*/
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	//切片操作
	s := "I am a student."
	m, n := 2, 4

	fmt.Printf("s[m:n]: %v\n", s[m:n]) //获取字符串s索引位置从m到n-1的值
	fmt.Printf("s[:n]: %v\n", s[:n])   //获取字符串s索引位置从0到n-1的值
	fmt.Printf("s[m:]: %v\n", s[m:])   //获取字符串s索引位置从m到len(s)-1的值
	fmt.Printf("s[:]: %v\n", s[:])     //获取字符串s
	fmt.Printf("s[m]: %v\n", s[m])     //获取字符串s索引位置m的字符的ascii值

	//分割字符串
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))

	//查询某个字符串是否包含指定的字符串
	fmt.Printf("strings.Contains(s, \"student\"): %v\n", strings.Contains(s, "student"))

	//判断前缀是否是指定的字符串
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))

	//判断后缀是否是指定的字符串
	fmt.Printf("strings.HasSuffix(s, \"student.\"): %v\n", strings.HasSuffix(s, "student."))

	//查找指定字符串第一次出现的索引位置
	fmt.Printf("strings.Index(s, \"a\"): %v\n", strings.Index(s, "aaa"))

	//查找指定字符串最后一次出现的索引位置
	fmt.Printf("strings.LastIndex(s, \"a\"): %v\n", strings.LastIndex(s, "a"))

	//拼接字符串数组
	fmt.Printf("strings.Join([]string{\"i\", \"am\", \"a\", \"student.\"}, \" \"): %v\n", strings.Join([]string{"i", "am", "a", "student."}, " "))
}
