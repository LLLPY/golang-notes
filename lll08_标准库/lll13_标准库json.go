package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person1 struct {
	Name   string
	Age    int
	PetDog Dog
}

type Dog struct {
	Name string
	Age  int
}

func main() {

	erha := Dog{Name: "二哈", Age: 3}
	tom := Person1{Name: "Tom", Age: 18, PetDog: erha}
	fmt.Printf("tom: %v\n", tom)

	//1.将一个struct实例转换成一个json对象
	//Marshal接收一个任何类型的对象，然后会返回对应json字符串的字节切片
	b, _ := json.Marshal(tom)
	json_b := string(b)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("json_b: %v\n", json_b)

	//2.将一个json字符串转成一个struct Unmarshal接收一个json字符串的字节切片和一个任意struct对象的地址
	//Unmarshal会将json字符串对象的值相应的赋给这个struct对象
	var tom2 Person1
	json.Unmarshal([]byte(json_b), &tom2)
	fmt.Printf("tom2: %v\n", tom2)

	//3.从io流中获取json字符串，然后转成struct
	f, _ := os.Open("demo.json")
	defer f.Close()
	json_decoder := json.NewDecoder(f)
	var tom3 map[string]interface{}
	json_decoder.Decode(&tom3)
	for k, v := range tom3 {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}

	//4.将struct实例转成json后写入文件
	jerry := Person1{Name: "Jerry", Age: 16, PetDog: erha}
	f2, _ := os.OpenFile("demo2.json", os.O_RDWR|os.O_CREATE, 0777)
	defer f2.Close()
	json_encoder := json.NewEncoder(f2)
	json_encoder.Encode(jerry)

	//从json文件中读取写入的内容
	res := make([]byte, 100)
	f3, _ := os.Open("demo2.json")
	f3.Read(res)
	defer f3.Close()
	fmt.Printf("res: %v\n", res)

}
