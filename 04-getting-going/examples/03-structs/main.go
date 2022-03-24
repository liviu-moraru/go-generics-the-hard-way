package main

import "fmt"

type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type SumFn[K Numeric] func(...K) K
type Ledger[T ~string, K Numeric] struct {
	//ID identifies the ledger
	ID T
	// Amounts is a list of monies associated with this ledger.
	Amounts []K
	// SumFn is a function that can be used to sum the amounts
	// in this ledger.
	SumFn SumFn[K]
}

func Sum[T Numeric](values ...T) T {
	var s = new(T)
	for i := 0; i < len(values); i++ {
		*s += values[i]
	}
	return *s
}

func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}

func main() {
	Ledger[string, int]{
		ID:      "acct-l",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	}.PrintIDAndSum()
}
