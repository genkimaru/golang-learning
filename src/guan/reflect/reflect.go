package main

import (

	"fmt"
	 "reflect"
	 "guan/pkg"
)

type  Person struct {

	Name string
	Age int

}

func (person Person) toString() {
	fmt.Printf(" person is %v \n"  , person)

}

func main(){


	mypkg.Add()

	print := fmt.Printf
	p := Person{"jack" , 20}
	p.toString()
	v := reflect.ValueOf(&p)
	print("v type is %T \n" , v)
	age := v.Elem().FieldByName("Age")
	// age := v.Elem().Field(1)
	print("age type is %T \n" , age)

	if age.IsValid() {
		// A Value can be changed only if it is 
		// addressable and was not obtained by 
		// the use of unexported struct fields.
		if age.CanSet() {
			// change value of N
			if age.Kind() == reflect.Int {
				x := int64(7)
				if !age.OverflowInt(x) {
					age.SetInt(x)
				}
			}
		}
	}

	p.toString()
}