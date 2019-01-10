package main

import (
	"fmt"

	"guan/reflect/example3/tools"
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

	tools.SetValueForField(&s, "MaxHeight", 100)

	fmt.Printf(" %v ", s)

	result := tools.CallMethod(&s, "Add", 100, 200)
	fmt.Println(result)
	//result.(int[])
}
