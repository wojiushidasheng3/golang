package main

import "fmt"

func testgo() {
	var originalCount = 10
	var eatenCount = 4

	fmt.Println("I started with ", originalCount, "apples")
	fmt.Println("Some jerk ate", eatenCount, " apples")
	fmt.Println("There are ", originalCount-eatenCount, "apples left")
}
