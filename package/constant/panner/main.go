package main

import (
	"fmt"
	"hellogo/package/constant/dates"
)

func main() {
	days := 3
	fmt.Println("你当前预约在", days, "天之后")
	fmt.Println("你的续约在", dates.DaysInWeek, "天之后")
}
