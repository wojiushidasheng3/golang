package main

import (
	"fmt"
	"math"
)

func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("输入数字无效")
	}
	number = math.Sqrt(number)
	return number, nil
}

func main() {
	num, err := Sqrt(3.5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("平方根为%0.2f", num)
}
