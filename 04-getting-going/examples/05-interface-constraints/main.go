package main

import (
	"fmt"
)

type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type SumFn[K Numeric] func(...K) K

type Ledgerish[T ~string, K Numeric] interface {
	~struct {
		ID      T
		Amounts []K
		SumFn   SumFn[K]
	}
	PrintIDAndSum()
}

type Ledger[T ~string, K Numeric] struct {
	//ID identifies the ledger
	ID T
	// Amounts is a list of monies associated with this ledger.
	Amounts []K
	// SumFn is a function that can be used to sum the amounts
	// in this ledger.
	SumFn SumFn[K]
}

func Sum[K Numeric](values ...K) K {
	var sum K
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}

	return sum

}

func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}

// PrintLedger emits a ledger's ID and total amount on a single line
// to stdout.

func PrintLedger[T ~string, K Numeric, L Ledgerish[T, K]](l L) {
	l.PrintIDAndSum()
}

func PrintL[T ~string, K Numeric](l Ledger[T, K]) {
	l.PrintIDAndSum()
}

func main() {
	PrintLedger(Ledger[string, complex64]{
		ID:      "fake",
		Amounts: []complex64{1, 2, 30},
		SumFn:   Sum[complex64],
	})

}
