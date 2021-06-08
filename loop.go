package main

import (
	"bufio"
	"fmt"
	"os"
)

func printFileFun(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// 我们针对文本内容做处理
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	printFileFun("abc.txt")
	//forever()
}
