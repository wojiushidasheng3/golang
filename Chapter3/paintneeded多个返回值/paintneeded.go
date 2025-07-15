package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("width %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("height %0.2f is iinvalid", height)
	}

	area := width * height
	return area, nil
}

func main() {
	amount, err := paintNeeded(3.5, 2.4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}
}
