package main

import "fmt"

//求平均值的闭包
func make_avger() func(int) float64 {

	s := make([]int, 0) //切片

	return func(num int) float64 {
		s = append(s, num)
		sum := 0
		for _, v := range s {
			sum += v
		}
		return float64(sum / len(s))
	}

}
func main() {

	avg := make_avger()
	for i := 0; i < 10; i++ {
		fmt.Printf("%.2f\n", avg(i))

	}
}
