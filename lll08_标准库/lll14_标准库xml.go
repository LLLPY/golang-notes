package main

import (
	"encoding/xml"
	"fmt"
)

type Person2 struct {
	XMLName xml.Name `xml:"Person"` //xml中的根元素
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func main() {

	p := Person2{
		Name:  "Tom",
		Age:   18,
		Email: "Tom@110.com",
	}
	//1.将一个struct实例转成一个xml
	// xml.Marshal(&p) //这个是没有缩进的
	b, _ := xml.MarshalIndent(&p, "", " ")
	fmt.Printf("%v\n", string(b))

	//2.将一个xml转成一个struct
	var p2 Person2
	xml.Unmarshal(b, &p2)
	fmt.Printf("p2: %v\n", p2)

}
