package main

import "fmt"

func SayHi() {
	fmt.Println("hello world")
}

func result(line string, times int) {
	for i := 0; i < 3; i++ {
		fmt.Println(line)
	}
}

func main() {
	result("hello", 3)
}
