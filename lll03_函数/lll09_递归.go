package main

//使用递归实现斐波那契数列
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}

func main() {
	res := fib(10)
	print(res)
}
