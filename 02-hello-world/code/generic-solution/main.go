package main

import (
	"fmt"
)

type request struct {
	host *string
	port *int
}

func Ptr[T any](value T) *T {
	return &value
}

func print(r request) {
	fmt.Print("request: host=")
	if r.host != nil {
		fmt.Print(*r.host)
	}
	fmt.Print(", port=")
	if r.port != nil {
		fmt.Printf("%d", *r.port)
	}
	fmt.Println()
}

func main() {

	print(request{
		host: Ptr("local"), // needs a *string
		port: Ptr(8080),    // needs a *int
	})
}
