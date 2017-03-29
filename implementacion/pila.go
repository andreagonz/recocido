package recocido

type Pila []int

func (s Pila) Push(v int) Pila {
    return append(s, v)
}

func (s Pila) Pop() (Pila, int) {
	l := len(s)
	if l == 0 {
		return nil, 0
	}
	return  s[:l-1], s[l-1]
}
