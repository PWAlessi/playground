package main

import "fmt"

func main() {
	foo()
}

func foo() {
	fmt.Println("Foo")
	bar()
}

var bar = func() {
	fmt.Println("bar")
}
