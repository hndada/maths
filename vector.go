package maths

// Defining a Vector as a slice of any type is not proper;
// It would allow random types in one vector, e.g. Vector{1, "hello", 3.14}
type Vector[T any] []T
