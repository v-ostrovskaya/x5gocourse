package main

import "fmt"

func main() {
	orders := []int{12000, 7500, 31000, 4500}

	if len(orders) == 0 {
		fmt.Println("Нет заказов: статистику рассчитать нельзя")
		return
	}

	minOrder := orders[0]
	maxOrder := orders[0]
	total := 0

	for _, amount := range orders {
		if amount < minOrder {
			minOrder = amount
		}

		if amount > maxOrder {
			maxOrder = amount
		}

		total += amount
	}

	average := float64(total) / float64(len(orders))

	fmt.Println("Минимум:", minOrder)
	fmt.Println("Максимум:", maxOrder)
	fmt.Println("Сумма:", total)
	fmt.Println("Среднее:", average)
}
