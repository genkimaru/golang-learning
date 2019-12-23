package main

import (
	"fmt"

	"github.com/codegangsta/inject"
)

func main() {

	injector := inject.New()

	fmt.Println(injector)
}
