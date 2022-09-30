package main

func init_var() int {
	print("初始化变量...\n")
	return 10
}

func init() {
	print("init函数1...\n")
}

func init() {
	print("init函数2...\n")
}

var i int = init_var()

func main() {
	print("main函数被执行了...\n")

}
