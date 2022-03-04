package main

import (
	"fmt"

	"github.com/genkimaru/golang-learning/pointer/pointer"
)

func main() {

	a := 1
	fmt.Printf("a address %p\n", &a)

	pointer.Zero(a)

	fmt.Println("after zero , a = ", a)

	pointer.ZeroPointer(&a)

	fmt.Println("after zeropointer , a = ", a)

	person := pointer.Person{Name: "wang", Age: 20}

	fmt.Println(person.SetValue("guan", 30))

	person.SetValue2("guan2", 40)

	fmt.Println(" person      ", person)

	personAddr := &person

	fmt.Printf(" personAddr type %T \n", personAddr)

	fmt.Println(" ------ ", personAddr.SetValue("guan", 30))

	// fmt.Println(" ------ ", personAddr.SetValue2("guan2", 34))

}
