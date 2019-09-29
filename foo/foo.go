package foo

import "fmt"

// Foo calls Bar.
func Foo() {
	if mock.Foo != nil {
		mock.Foo()
		return
	}
	Bar()
}

// Bar prints "Bar!".
func Bar() {
	if mock.Bar != nil {
		mock.Bar()
		return
	}
	fmt.Println("Bar!")
}

var mock *MockHandler = nil

// MockHandler handles mock functions.
type MockHandler struct {
	Foo func()
	Bar func()
}

// Mock mocks module functions.
func Mock(m *MockHandler) {
	mock = m
}
