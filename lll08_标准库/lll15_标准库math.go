package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	//常量
	fmt.Printf("math.MaxFloat64: %v\n", math.MaxFloat64)
	fmt.Printf("math.MaxFloat32: %v\n", math.MaxFloat32)
	fmt.Printf("math.MaxInt64: %v\n", math.MaxInt64)
	fmt.Printf("math.MaxInt32: %v\n", math.MaxInt32)
	fmt.Printf("math.MaxInt16: %v\n", math.MaxInt16)
	fmt.Printf("math.MaxInt8: %v\n", math.MaxInt8)
	fmt.Printf("math.MaxInt: %v\n", math.MaxInt)
	fmt.Printf("math.Pi: %.20f\n", math.Pi)

	//数学函数

	//绝对值
	fmt.Printf("math.Abs(-100): %v\n", math.Abs(-100))

	//幂函数
	fmt.Printf("math.Pow(2, 3): %v\n", math.Pow(2, 3)) //2^3

	//开平方
	fmt.Printf("math.Sqrt(4): %v\n", math.Sqrt(4))

	//开立方
	fmt.Printf("math.Cbrt(8): %v\n", math.Cbrt(8))

	//向上取整
	fmt.Printf("math.Ceil(10.2): %v\n", math.Ceil(10.2))
	//向下取整
	fmt.Printf("math.Floor(10.8): %v\n", math.Floor(10.8))
	//四舍五入
	fmt.Printf("math.Round(10.4): %v\n", math.Round(10.4))

	//取余
	fmt.Printf("math.Mod(10, 3): %v\n", math.Mod(10, 3))

	//三角函数
	angle := math.Pi / 180 //1度
	fmt.Printf("math.Sin(90): %v\n", math.Sin(90*angle))
	fmt.Printf("math.Cos(60): %v\n", math.Cos(60*angle))
	fmt.Printf("math.Tan(90): %v\n", math.Tan(45*angle))

	//随机数
	rand.Seed(time.Now().UnixMicro()) //设置随机数种子
	for i := 0; i < 5; i++ {
		fmt.Printf("rand.Intn(100): %v\n", rand.Intn(100)) //100以内的随机整数
		fmt.Printf("rand.Float64(): %v\n", rand.Float64()) //0到1的随机小数

	}

}
