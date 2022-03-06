// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java and Ruby and
// the overloaded single result / error value sometimes
// used in C. Go's approach makes it easy to see which
// functions return errors and to handle them using the
// same language constructs employed for any other,
// non-error tasks.

package main

import (
	"errors"
	"fmt"
)

// It's possible to use custom types as `error`s by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.
type argError struct {
	arg  int
	prob string
}

// ArgError doc
type ArgError struct {
	Arg  int
	Prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func (e ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.Prob)
}

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return -1, errors.New("can't work with 42")

	}

	// A `nil` value in the error position indicates that
	// there was no error.
	return arg + 3, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// In this case we use `&argError` syntax to build
		// a new struct, supplying values for the two
		// fields `arg` and `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func f3(arg int) (int, error) {
	if arg == 42 {
		return -1, ArgError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// error is an interface
func f4(arg int) ArgError {

	return ArgError{arg, "can't work with it"}
}

func main() {

	// The two loops below test out each of our
	// error-returning functions. Note that the use of an
	// inline error check on the `if` line is a common
	// idiom in Go code.
	// for _, i := range []int{7, 42} {
	// 	if r, e := f1(i); e != nil {
	// 		fmt.Println("f1 failed:", e)
	// 	} else {
	// 		fmt.Println("f1 worked:", r)
	// 	}
	// }
	// for _, i := range []int{7, 42} {
	// 	if r, e := f2(i); e != nil {
	// 		fmt.Println("f2 failed:", e)
	// 	} else {
	// 		fmt.Println("f2 worked:", r)
	// 	}
	// }

	_, e := f3(42)
	fmt.Printf("-------%T\n", e)

	e2 := f4(11)
	fmt.Println(e2.Prob)
	// ee := ArgError{11, "ee msg"}
	// fmt.Printf("-------%T\n", &ee)
	// fmt.Println(e.Prob)
	// fmt.Println(*e.prob)
	// if ae, ok := e.(*ArgError); ok {
	// 	fmt.Println(ok)
	// 	fmt.Println(ae.Arg)
	// 	fmt.Println(ae.Prob)
	// }
}
