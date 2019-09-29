package main

import (
	"mocktestapp/foo"
)

func main() {
	foo.Mock(&foo.MockHandler{})
	foo.Foo()
}
