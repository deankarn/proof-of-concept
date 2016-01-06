package main

import "fmt"

func main() {

	c := &MyContext{DefaultContext: DefaultContext{}}
	c.Reset("TEST")
	// c.DefaultContext.Reset("TEST")
	// c.Test()
	// c.DefaultContext.Test()
}

// Context desc
type Context interface {
	Test()
	Reset(val string)
}

// DefaultContext desc
type DefaultContext struct {
}

// Test ...
func (c *DefaultContext) Test() {
	fmt.Println("DefaultContext.Test()")
}

// Reset ...
func (c *DefaultContext) Reset(val string) {
	fmt.Println("DefaultContext.Reset(val string)")
}

// MyContext ...
type MyContext struct {
	DefaultContext
}

// Reset ...
func (c *MyContext) Reset(val string) {
	fmt.Println("MyContext.Reset(val string)")
	c.DefaultContext.Reset(val)
}

var _ Context = &MyContext{}

// var _ Context = &MyContext{DefaultContext{}}
