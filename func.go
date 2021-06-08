package main

import (
	"fmt"
	"math"
)

func div(a, b int) (q, r int) {
	return a / b, a % b
}
func eval(a, b int, op string) (int, error) {
	switch op {
	case " + ":
		return a + b, nil
	case " - ":
		return a - b, nil
	case " * ":
		return a * b, nil
	case " % ":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupoorted operator " + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling")
	return op(a, b)
}

func main() {
	//fmt.Println(eval(13, 4, "x"))
	if result, err := eval(13, 4, "x"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}
	q, _ := div(13, 3)
	fmt.Println(q)

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))
}
