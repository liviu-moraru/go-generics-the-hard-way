package main

import (
	"fmt"
)

type request struct {
	host *string
	port *int
}

func PtrInt(i int) *int {
	return &i
}

func PtrStr(s string) *string {
	return &s
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
		host: PtrStr("local"), // needs a *string
		port: PtrInt(80),      // needs a *int
	})
}
