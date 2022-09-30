package main

import (
	"fmt"
)

func main() {

	a := 100
	if a > 10 {
		fmt.Println("hello")
	}

	if b := 100; b > 10 {
		fmt.Println("if条件成立！")
		fmt.Printf("b: %v\n", b)
	} else {
		fmt.Println("if条件不成立！")
		fmt.Printf("b: %v\n", b)
	}
	// fmt.Printf("b: %v\n", b)

	if age := 18; age > 18 {
		fmt.Println("你已经成年了！")
	} else {
		fmt.Println("你还未成年！")
	}

	var aa int
	fmt.Println("请输入一个数：")
	fmt.Scan(&aa)
	if aa == 100 {
		fmt.Printf("aa: %v\n", aa)
	} else if aa < 100 {
		fmt.Println("太小了！")
	} else {
		fmt.Println("太大了！")
	}

	//switch语句
	aaa := 1
	switch aaa {
	case 1:
		fmt.Println("唱歌！")
	case 2:
		fmt.Println("跳舞！")
	case 3:
		fmt.Println("唱歌+跳舞！")
	default:
		fmt.Println("才艺表演。")
	}

	var day int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&day)
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日！")
	case 6, 7:
		fmt.Println("假期！")
	default:
		fmt.Println("输入有误！")
	}

	//case后接条件表达式
	var n int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&n)
	switch {
	case n == 10:
		fmt.Println("猜对了！")
	case n < 10:
		fmt.Println("太小了！")
	case n > 10:
		fmt.Println("太大了！")
	default:
		fmt.Println("不可能执行到这条语句，除非上面的case都不成立。")
	}

	//fallthrough

	bb := 100
	switch bb {
	case 100:
		fmt.Printf("bb: %v\n", bb)
		fallthrough
	case 200:
		fmt.Println("200")
		fallthrough
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("case都不成立就执行这里。")
	}

}
