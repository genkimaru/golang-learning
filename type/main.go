package main

import "fmt"

// 这个是new type的 definition
type MyFloat float64

// 这个是 map[string]interface{} 类型的alias
type myMap = map[string]interface{}

func (myFloat MyFloat) M() MyFloat {
	return myFloat * myFloat
}

func (f MyFloat) String() string {
	return fmt.Sprintf("this is my float : %f ", f)
}

func main() {

	var f MyFloat = 2.0
	fmt.Println(f.M())

	m := make(myMap)
	m["a"] = 1

	m2 := myMap{
		"a": 1,
		"b": 2,
	}

	delete(m2, "a")

	v, ispresent := m2["a"]
	fmt.Println("a is ", v, " , is present :", ispresent)

}
