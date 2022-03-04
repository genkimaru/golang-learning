package pointer

import "fmt"

//Person struct
type Person struct {
	Name string
	Age  int
}

// Zero 值传递
func Zero(i int) {
	fmt.Printf("in zero function  , i address  %p\n", &i)
	i = 0
}

// ZeroPointer 接受一个指针地址类型
func ZeroPointer(i *int) {
	fmt.Printf("in zeropointer function  , i address  %p\n", i)
	*i = 0
}


//SetValue doc
func (person Person) SetValue(name string, age int) Person {

	person.Name = name
	person.Age = age
	return person
}


//SetValue2 doc
func (person *Person) SetValue2(name string, age int) {

	(*person).Name = name
	(*person).Age = age
}
