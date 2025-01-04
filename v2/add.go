// Package adder includes a math function to add two integers
package adder

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes in two numbers and returns the result, these numbers can be either integers or floats
// To learn more about maths check this [site]: https://www.mathisfun.com/numbers/addition.html

func Add[T Number](a, b T) T {
	return a + b
}
