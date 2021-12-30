package test

import (
	"fmt"
	"testing"
)

// Client1 Experiment.
type Client1 struct {
}

type Opt1 func(*Client1) error

// Abc The return type of the function is another function of type Opt1
func Abc() Opt1 {
	f := func(c *Client1) error {
		return nil
	}
	return f
}

// NewClientWithOpts1 The function accepts types of function Opt1
func NewClientWithOpts1(ops ...Opt1) (*Client1, error) {
	return &Client1{}, nil
}
func FromEnv1(c *Client1) error {
	return nil
}

func FromEnv2(c *Client1) int {
	return 10
}

func TestFunctionType(t *testing.T) {
	var f1 Opt1
	f1 = FromEnv1
	fmt.Printf("%T\n", f1)

	var f2 Opt1
	f2 = Abc()
	fmt.Printf("%T\n", f2)
}
