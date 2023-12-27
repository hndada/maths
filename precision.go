package maths

var (
	EqualEpsilon         = 1e-6
	LimitEpsilon         = 1e-20
	DifferentiateEpsilon = 1e-6
	IntegrateEpsilon     = .5 * 1e-4
)

// less than 0.00001 difference is considered equal
func isEqual(a, b Real) bool {
	h := EqualEpsilon
	return a-h < b && b < a+h
}
