package main

import (
	"errors"
	"fmt"
	"time"
)

//自定义errors
type MyError struct {
	When time.Time
	What string
}

/*
type error interface {
	Error() string
}


error是一个接口，只要实现了Error方法，就可以是一个error

*/

func (e MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}

//检测字符串是否为空
func check_str(s string) (err error) {

	if s == "" {
		err = errors.New("字符串不能为空...")
	} else {
		err = MyError{
			When: time.Date(2022, 11, 11, 11, 11, 11, 11, time.UTC),
			What: fmt.Sprintf("%v 不是一个空字符串...", s),
		}
	}
	return

}
func main() {

	s := ""
	err := check_str(s)
	fmt.Printf("err: %v\n", err)
	s2 := "hello python"
	err2 := check_str(s2)
	fmt.Printf("err2: %v\n", err2)

}
