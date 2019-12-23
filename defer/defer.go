package main

import (
	"errors"
	"fmt"
)

//main This will be executed at the end of the enclosing function (main),

func main() {
	defer beforeCloseMain()
	deferTest()
}

func deferTest() {

	fmt.Println("  start... ")
	defer close()
	var err = errors.New("can't work with 42")
	if err != nil {
		panic(err)
	} else {

		fmt.Println("working  ..")
	}

}

func close() {

	fmt.Println("close deferTest")
}

func beforeCloseMain() {
	fmt.Println("close main")
}
