package maths

type Element any

type Set interface {
	Contains(Element) bool
}

type Mapping struct {
	Domain   Set
	Codomain Set
	Map      func(Element) Element
}
