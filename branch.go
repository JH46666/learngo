package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

//func eval(a, b int, op string) {
//	var result int
//	switch op {
//	case " + ":
//		result = a + b
//	case " - ":
//		result = a - b
//	case " * ":
//		result = a * b
//	case " % ":
//		result = a / b
//	default:
//		painc("unsupoorted operator " + op)
//	}
//	return result
//}

func getScore(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	const filename = "abc.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
	//eval()
	fmt.Println(
		getScore(100),
	)
	fmt.Println(
		converToBin(100),
		converToBin(13), converToBin(0),
	)
}
