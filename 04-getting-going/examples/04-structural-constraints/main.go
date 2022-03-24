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

type Ledger[T ~string, K Numeric] struct {
	//ID identifies the ledger
	ID T
	// Amounts is a list of monies associated with this ledger.
	Amounts []K
	// SumFn is a function that can be used to sum the amounts
	// in this ledger.
	SumFn SumFn[K]
}

type CustomLedger struct {
	ID      ID
	Amounts []uint64
	SumFn   SumFn[uint64]
}

type LedgerNode struct {
	ID      string
	Amounts []uint64
	SumFn   SumFn[uint64]
	Next    *LedgerNode
}

type ID string

func SomeFunc[
	T ~string,
	K Numeric,
	L ~struct {
		ID T
		// Amounts is a list of monies associated with this ledger.
		Amounts []K
		// SumFn is a function that can be used to sum the amounts
		// in this ledger.
		SumFn SumFn[K]
	},
](l L) {

	//The fields are not accessible
	/*
		https://github.com/golang/go/issues/48522
		"We don't currently support field accesses of this kind even though the proposal says that this should/could work. We may not support this for Go1.18 as it doesn't seem like an essential feature. There's a trivial work-around that uses a method:"

		The workaround: use interface constraint with struct and method definition
	*/

	//the following line does not compile
	//fmt.Println(l.ID)

	fmt.Printf("%+v", l)
}

func Sum[T Numeric](values ...T) T {
	var s = new(T)
	for i := 0; i < len(values); i++ {
		*s += values[i]
	}
	return *s
}

func main() {
	SomeFunc(Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	})

	SomeFunc[ID, float32, struct {
		ID      ID
		Amounts []float32
		SumFn   SumFn[float32]
	}](struct {
		ID      ID
		Amounts []float32
		SumFn   SumFn[float32]
	}{
		ID:      ID("fake"),
		Amounts: []float32{1, 2, 3},
		SumFn:   Sum[float32],
	})
	SomeFunc(struct {
		ID      ID
		Amounts []float32
		SumFn   SumFn[float32]
	}{
		ID:      ID("fake"),
		Amounts: []float32{1, 2, 3},
		SumFn:   Sum[float32],
	})
	SomeFunc(CustomLedger{
		ID:      ID("new fake"),
		Amounts: []uint64{100, 200},
		SumFn:   Sum[uint64],
	})

	//does not compile. Structural constraints must match the struct exactly
	/* SomeFunc(LedgerNode{
		ID:      ID("new fake"),
		Amounts: []uint64{100, 200},
		SumFn:   Sum[uint64],
	}) */

}
