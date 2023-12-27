package maths

// Definition: A function from the natural numbers to the reals.
type Sequence interface {
	Term(int) Real
}

func SumSequence(s Sequence, n int) Real {
	var sum Real
	for i := 0; i < n; i++ {
		sum += s.Term(i)
	}
	return sum
}

type Recurrence struct {
	Firsts []Real
	Next   func(...Real) Real
	cached map[int]Real
}

func (r *Recurrence) Term(n int) Real {
	if r.cached == nil {
		r.cached = make(map[int]Real)
	}
	if n < len(r.Firsts) {
		return r.Firsts[n]
	}
	if v, ok := r.cached[n]; ok {
		return v
	}
	// Create a slice to hold arguments for r.Next
	args := make([]Real, 0, len(r.Firsts))
	for i := 0; i < len(r.Firsts); i++ {
		args = append(args, r.Term(n-1-i))
	}
	v := r.Next(args...)
	r.cached[n] = v
	return v
}
