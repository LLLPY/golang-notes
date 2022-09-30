package main

import "fmt"

func main() {

	var p *int
	var pp **int
	var ppp ***int
	a := 10
	p = &a    //将指针p指向a的地址 即*p=a=10
	pp = &p   //将指针pp指向p的地址 即*pp=p-> *(*pp)=*p=a=10
	ppp = &pp //将指针ppp指向pp的地址 即*ppp=pp-> *(*ppp)=*pp -> *(*(*ppp))=*(*pp)=*p=a=10

	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %v\n", &p)
	fmt.Printf("pp: %v\n", pp)
	fmt.Printf("ppp: %v\n", ***ppp)
	fmt.Println(*p == a)
	fmt.Println(*(*pp) == a)
	fmt.Println(*(*(*ppp)) == a)

}
