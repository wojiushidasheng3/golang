package main

import (
	"fmt"
	"math"
)

func floatResult(number float64) (int, float64) { //这里要求return的类型分别是int和folat64
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	cans, remainder := floatResult(1.26)
	fmt.Println(cans, remainder)
}

/*
func main() { // 用对应的类型变量接收返回值
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
}

// manyReturns 函数返回三个值，分别为 int、bool 和 string 类型
func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

*/
