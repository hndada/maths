package main

import (
	"fmt"

	"github.com/hndada/maths"
)

type (
	Real         = maths.Real
	RealFunction = maths.RealFunction
)

func main() {
	fn := func(x Real) Real { return x * x }
	f := RealFunction{Evaluate: fn}
	fmt.Println(f.Evaluate(2))
	fmt.Println(f.Limit(2))
	fmt.Println(f.IsContinuous(2))
	fmt.Println(f.Differentiate(2))
	fmt.Println(f.Integrate(0, 2))
}
