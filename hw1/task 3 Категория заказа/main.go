package main

import "fmt"

func main() {
	amount := 25000

	if amount < 0 {
		fmt.Println("Error: order amount cannot be negative")
		return
	}

	switch {
	case amount < 10_000:
		fmt.Println("small")
	case amount < 50_000:
		fmt.Println("regular")
	default:
		fmt.Println("large")
	}
}
