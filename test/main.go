package main

import (
	"fmt"

	. "github.com/hndada/maths"
)

func main() {
	f := Function{Evaluate: func(x Real) Real {
		return x * x
	}}
	fmt.Println(f.Evaluate(2))
	fmt.Println(f.Limit(2))
	fmt.Println(f.IsContinuous(2))
	fmt.Println(f.Differentiate(2))
	fmt.Println(f.Integrate(0, 2))
}
