package main

import (
	"fmt"

	"github.com/genkimaru/golang-learning/interface/itf"
)

func main() {
	s := itf.FetchStudent()
	fmt.Printf("---- %T \n", s)
	// 通过s 来获得age和class
	ss := s.(itf.Student) // type conversion
	fmt.Println("----", ss.Age)

	w := itf.FetchWorker()
	fmt.Printf("---- %T \n", w)
	ww := w.(*itf.Worker) // type conversion
	fmt.Println("----", ww.Age)
}
