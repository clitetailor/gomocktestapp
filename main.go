package mocktestapp

import "fmt"

func Foo() {
	if mock.Foo != nil {
		mock.Foo()
		return
	}
	bar()
}

func bar() {
	if mock.Bar != nil {
		mock.Bar()
		return
	}
	fmt.Println("Bar!")
}

var mock *MockHandler = nil

type MockHandler struct {
	Foo func()
	Bar func()
}

func Mock(m *MockHandler) {
	mock = m
}

func main() {
	Foo()
}
