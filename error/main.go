package main

import (
	"fmt"
	"math"
)

// type ErrNegativeSqrt float64

type NegativeSqrtErr struct {
	reason string
	number float64
}

func (e *NegativeSqrtErr) Error() string {
	return fmt.Sprintf("reason : %s  , number : %f", e.reason, e.number)

}

func NegativeSqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	} else {
		return x, &NegativeSqrtErr{"can't use negtive value", x}
	}
}

func Sqrt(x float64) (float64, error) {

	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	//	fmt.Println(Sqrt(-2))
	fmt.Println(NegativeSqrt(2))
	fmt.Println(NegativeSqrt(-2))

}
