package user

import "fmt"

func Hello() string {
	return "hello world!"
}

func hello() string {
	return "hello world!"
}

func main() {
	s := Hello()
	fmt.Println(s)
}
