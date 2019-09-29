package foo

import (
	"fmt"
)

type Foo struct {
	foo int
	bar int
}

type Impl interface {
	CallFoo()
	CallBar()
}

func NewFoo(foo int, bar int) *Foo {
	return &Foo{
		foo: foo,
		bar: bar,
	}
}

func (foo *Foo) CallFoo() {
	fmt.Println(foo.foo)

	foo.CallBar()
}

func (foo *Foo) CallBar() {
	fmt.Println(foo.bar)
}

var FooMod Impl = NewFoo(3, 4)
