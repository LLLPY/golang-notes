package main

import "fmt"

func main() {

	var m = map[string]string{
		"name": "Tom",
		"age":  "18",
		"num":  "1234",
	}

	//for range遍历
	//1.只拿到key
	for k := range m {
		fmt.Printf("k: %v v: %v\n", k, m[k])
	}

	//2.同时拿到k和v
	for k, v := range m {
		fmt.Printf("k: %v v: %v\n", k, v)
	}

}
