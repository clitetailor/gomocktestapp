package foo_test

import (
	"mocktestapp/foo"
	"testing"

	"github.com/stretchr/testify/mock"
)

type FooMock struct {
	mock.Mock
	foo.Foo
}

func (f *FooMock) CallFoo() {
	f.Called()
	f.Foo.CallFoo()
}

func (f *FooMock) CallBar() {
	f.Called()
}

func TestFoo(t *testing.T) {
	f := &FooMock{}

	foo.FooMod = f

	f.On("CallBar").Times(2).Return()

	f.CallFoo()
	f.CallFoo()

	f.AssertExpectations(t)
}
