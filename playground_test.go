package main

import (
	"fmt"
	"testing"
)

func Test_foo(t *testing.T) {
	foo()
}

func Test_bar(t *testing.T) {
	bar()
}

func Test_mockBar(t *testing.T) {
	// Here I want to test foo but mock bar
	// I want a different bar() func that prints RAB instead
	rab := func() {
		fmt.Println("RAB")
	}

	// Get a function pointer to bar
	bar = rab

	// Should print foo RAB
	foo()
}
