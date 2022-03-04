package main

import "fmt"

func main() {

	// 类型转换
	var a int = 10
	var b int64
	b = int64(a)
	fmt.Printf("---- %T \n", b)

	// 类型断言 type assertion
	var i interface{} = 11
	j := i.(int)
	fmt.Println("---- ", j)

	var x interface{} = "xxx"
	y, e := x.(string)
	if e {
		fmt.Println("---- ", y)
	}

}
