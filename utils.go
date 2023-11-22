package main

import "fmt"

type Case struct {
	name     string
	arg      any
	expected any
}

func (c *Case) errorInfo(actual any) string {
	return fmt.Sprintf("\nError: %s\nexpected: %v\nactual:   %v", c.name, c.expected, actual)
}
