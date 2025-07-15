package main

import "fmt"

func main() {
	//第一种循环
	fmt.Println("第一种")
	for x := 0; x < 3; x++ {
		fmt.Println(x)
	}

	//第二种循环
	fmt.Println("第二种")
	y := 1
	for y <= 3 {
		fmt.Println(y)
		y++
	}
}
