package main

import (
	"fmt"
)

/**

show the usage of fmt.Printf


https://www.jianshu.com/p/8be8d36e779c
*/

type Person struct {
	name    string
	age     int
	country Country
}

type Country struct {
	countrycode string
	countryname string
}

func main() {

	person := Person{"aaa", 12, Country{"usa", "America"}}
	fmt.Printf("------%v", person)
	fmt.Printf("------%+v", person)
	fmt.Printf("------%#v", person)

	i := 8

	fmt.Printf("integer usage , binary print , i = %b \n", i)
	fmt.Printf("integer usage , octonary print , i = %o \n", i)
	fmt.Printf("integer usage , hex print , i = %#x \n", i)
}
