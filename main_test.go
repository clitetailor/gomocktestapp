package mocktestapp_test

import (
	mta "mocktestapp"
	"testing"

	"github.com/stretchr/testify/mock"
)

var m = new(mock.Mock)

func BarMock() {
	m.Called()
}

func TestMain(t *testing.T) {
	mta.Mock(&mta.MockHandler{
		Bar: BarMock,
	})

	m.On("BarMock").Times(2).Return()

	mta.Foo()
	mta.Foo()

	m.AssertExpectations(t)

	mta.Mock(&mta.MockHandler{})
}
