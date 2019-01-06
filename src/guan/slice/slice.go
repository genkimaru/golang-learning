package main

import "fmt"

func main() {

	fmt.Println("-------")

	x := []byte{'r', 'o', 'a', 'd'}

	// x[5] = 'a'

	fmt.Printf("x %p \n", &(x[0]))

	x[0] = 'l'

	fmt.Printf("x %p \n", &(x[0]))

}
