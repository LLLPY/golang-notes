package main

import "fmt"

func main() {

	//for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\t", i)
	}
	fmt.Println("\n")
	//初始条件写在外面
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("j: %v\t", j)
	}

	fmt.Println("\n")
	//结束语句写在循环体内
	for i := 0; i < 10; {
		fmt.Printf("i: %v\t", i)
		i++

	}
	fmt.Println("\n")

	//永真循环
	k := 0
	for {
		fmt.Printf("k: %v ", k)
		k++
		if k == 10 {
			break
		}
	}
	fmt.Println("\n")
	//for range
	//遍历数组
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arr {
		fmt.Printf("%v:%v\n", i, v)

	}

	//遍历map
	var m = make(map[string]string, 0)
	m["name"] = "Tom"
	m["age"] = "18"
	m["num"] = "1234"
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)

	}

}
