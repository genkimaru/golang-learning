package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name      string
	MaxHeight int
}

func printInfo(i interface{}) {
	fmt.Printf(" type : %T \n address : %p \n value %v \n \n \n", i, &i, i)
}

func (s MyStruct) Add(a, b int) int {
	return a + b
}

func main() {

	var s = MyStruct{"aaa", 11}

	//fmt.Println(tools.ProduceStructTag(s, "json"))

	// printInfo(s)

	v := reflect.ValueOf(&s)

	// Kind 类型
	// Kind代表的是 所属的 静态类型
	// 比如  ：
	//reflect.int
	//reflect.String
	// reflect.Ptr 等等
	fmt.Println(v.Kind())

	// Type类型 代表所属的动态类型

	// 获取Type 类型有两种方式：
	// 一种是 通过 Value 结构体来获取
	fmt.Println(v.Type())

	// 第二种是 通过直接调用 reflect.TypeOf()的函数来获取。
	t := reflect.TypeOf(s)
	fmt.Println(t)

	// v := tools.New(s)

	// copyStruct := v.(*MyStruct)

	// printInfo(copyStruct)

	// fmt.Printf(" %p ", copyStruct)

	// var line string
	// line = fmt.Sprintf("{\n %s", "aaa")
	// line = fmt.Sprintf(" %s \n bbb", line)
	// line = fmt.Sprintf("  %s \n ccc \n }", line)
	// fmt.Println(line)
	// line := tools.PeekField(s)
	// fmt.Println(line)

	// line2 := tools.PeekMethod(s)
	// fmt.Println(line2)

	// tools.SetValueForField(&s, "MaxHeight", 100)

	// fmt.Printf(" %v ", s)

	// result := tools.CallMethod(&s, "Add", 100, 200)
	// fmt.Println(result)
	//result.(int[])
}
