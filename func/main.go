package main

import (
	"fmt"
	"math"
)

//接收一个函数，类型是 func(float64, float64) float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//golang的高阶函数。接收一个函数，类型是 func(float64, float64) float64 ,返回函数
func compute2(fn func(float64, float64) float64) func(float64) float64 {

	// 返回函数需要以匿名函数的方式
	return func(f float64) float64 {
		return f + fn(3, 4)
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(compute2(hypot)(10))
}
