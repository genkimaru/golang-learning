package main

import (

	"regexp"
	"fmt"
)

func main(){

	validPath := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

	path := "/edit/aaa"

	m := validPath.FindStringSubmatch(path)

	fmt.Println(m)
}
