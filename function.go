package maths

type Real = float64

type RealFunction struct {
	Evaluate func(Real) Real
}

type Function = RealFunction

func NewFunction(evaluate func(Real) Real) Function {
	return Function{Evaluate: evaluate}
}

func (f Function) Limit(x Real) (Real, bool) {
	h := LimitEpsilon
	left := f.Evaluate(x - h)
	right := f.Evaluate(x + h)
	if !isEqual(left, right) {
		return 0, false
	}
	return left, true
}

func (f Function) IsContinuous(x Real) bool {
	limit, exists := f.Limit(x)
	if !exists {
		return false
	}
	return f.Evaluate(x) == limit
}

func (f Function) Differentiate(x Real) Real {
	h := DifferentiateEpsilon
	return (f.Evaluate(x+h) - f.Evaluate(x)) / h
}

func (f Function) IsDifferentiable(x Real) bool {
	if !f.IsContinuous(x) {
		return false
	}
	return f.Differentiate(x) == f.Evaluate(x)
}

func (f Function) Integrate(a, b Real) Real {
	dx := IntegrateEpsilon
	var sum Real
	for x := a; x < b; x += dx {
		sum += f.Evaluate(x) * dx
	}
	return sum
}
