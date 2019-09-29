package foo_test

import (
	"mocktestapp/foo"
	"testing"

	"github.com/stretchr/testify/mock"
)

var m = new(mock.Mock)

func BarMock() {
	m.Called()
}

func TestMain(t *testing.T) {
	foo.Mock(&foo.MockHandler{
		Bar: BarMock,
	})

	m.On("BarMock").Times(2).Return()

	foo.Foo()
	foo.Foo()

	m.AssertExpectations(t)

	foo.Mock(&foo.MockHandler{})
}
