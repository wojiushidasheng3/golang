package main

import "fmt"

func PaintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10
}

func main() {
	var amount, total float64
	amount = PaintNeeded(3.4, 5.6)
	total += amount

	fmt.Printf("第一面墙的油漆用量为：%5.3f\n", amount)

	amount = PaintNeeded(5.6, 7.8)
	total += amount

	fmt.Printf("第二面墙的油漆用量为：%5.3f\n", amount)

	fmt.Printf("总油漆用量为: %5.3f", total)
}
