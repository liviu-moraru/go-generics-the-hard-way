package main

import "fmt"

// HasID is a structural constraint satisfied by structs with a single field
// called "ID" of type "string".
type HasID interface {
	~struct {
		ID string
	}
}

// CanGetID is an interface constraint satisfied by a type that has a function
// with the signature "GetID() string".
type CanGetID interface {
	GetID() string
}

// Unique satisfies the structural constraint "HasID" *and* the interface
// constraint "CanGetID."
type Unique struct {
	ID string
}

// UniqueName does *not* satisfiy the structural constraint "HasID," because
// while UniqueName has the field "ID string," the type also contains the field
// "Name string."
//
// Structural constraints must match *exactly*.
//
// UniqueName *does* satisfy the interface constraint "CanGetID."
type UniqueName struct {
	//Embedded struct
	Unique
	Name string
}

func (u UniqueName) GetID() string {
	return u.ID
}

// NewHasT returns a new instance of T.
func NewHasT[T HasID]() T {
	// Declare a new instance of T on the stack.
	var t T

	// Return the new T by value.
	return t
}

func NewT[T CanGetID]() T {
	var t T
	return t
}

// CanSetID is an interface constraint satisfied by a type that has a function
// with the signature "SetID(string)".
type CanSetID interface {
	SetID(string)
}

func (u *UniqueName) SetID(s string) {
	u.ID = s
}

// NewCanSetT returns a new instance of T.
func NewCanSetT[T CanSetID]() T {
	// Declare a new instance of T. Because T is constrained to be a
	// concrete type, it can easily be declared on the stack.
	var t T

	// Return the new T by value.
	return t
}

func main() {
	fmt.Printf("%T\n", NewHasT[Unique]())

	un := UniqueName{
		Unique: Unique{ID: "fake"},
		Name:   "XXX",
	}
	fmt.Printf("%+v. ID:%s\n", un, un.ID)
	// Next line doesn't compile. Error: UniqueName does not implement HasID
	// fmt.Printf("%T\n", NewHasT[UniqueName]())

	fmt.Printf("%T\n", NewT[UniqueName]())
	fmt.Printf("%T\n", NewCanSetT[*UniqueName]())
	// Print true
	fmt.Printf("%v\n", NewCanSetT[*UniqueName]() == nil)
}
