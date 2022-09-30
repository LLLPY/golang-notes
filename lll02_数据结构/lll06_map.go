package main

import (
	"fmt"
)

func main() {

	//map的定义
	//方法一
	var m1 map[string]string
	fmt.Printf("m1: %v\n", m1)

	//方法二
	m2 := make(map[string]string)
	fmt.Printf("m2: %v\n", m2)

	//map的初始化
	var m3 = map[string]string{
		"name": "Tom",
		"age":  "18",
	}
	fmt.Printf("m3: %v\n", m3)

	//增加/修改map
	m3["num"] = "1234"
	m3["age"] = "20"
	fmt.Printf("m3: %v\n", m3)

	//根据k获取v
	fmt.Printf("m3[\"name\"]: %v\n", m3["name"])

	//判断某个k是否存在  v,ok=m[k]--->如果k存在，ok为true，否则为false
	v, ok := m3["name"]
	if ok {
		print(v)
	}

	//删除某个k
	delete(m3, "age")
	fmt.Printf("m3: %v\n", m3)
}
