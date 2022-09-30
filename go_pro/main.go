package main

import (
	"fmt"
	"go_pro/user"
)

func main() {

	s := user.Hello()
	fmt.Println(s)

	var s2 = user.hello()
	fmt.Printf("s2: %v\n", s2)

}
