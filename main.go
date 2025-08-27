package main

import (
	"fmt"
	"strings"
)

func countDelivered(sliceOrderStatuses []string) int {

	count := 0

	for _, status := range sliceOrderStatuses {
		if strings.EqualFold(status, "delivered") {
			count++
		} else {
			continue
		}
	}

	fmt.Printf("Количество заказов со статусом 'delivered': %d\n", count)
	return count
}

func markCancelled(status *string) []string {

	if *status == "created" {
		*status = "cancelled"
	}
	return nil
}

func main() {

	var (
		orderID      int    = 1
		customerName string = "ИВАН"
		isDelivered  bool   = false
	)

	orderIDs := []int{}
	orderIDs = append(
		orderIDs,
		0,
		1,
		2,
	)

	orderStatuses := []string{
		"created", "shipped", "delivered", "cancelled", "delivered", "created",
	}

	orderCount := map[string]int{
		"Client1": 1,
		"Client2": 5,
	}

	isReadyToShip := orderCount["Client2"] > 2

	fmt.Println(
		orderID,
		customerName,
		isDelivered,
	)

	fmt.Println(
		len(orderIDs),
		cap(orderIDs),
	)

	fmt.Println(isReadyToShip)

	countDelivered(orderStatuses)

	fmt.Println("слайс статусов ДО отмены -", orderStatuses)

	for i := range orderStatuses {
		markCancelled(&orderStatuses[i])
	}

	fmt.Println("слайс статусов после отмены -", orderStatuses, "\n")

	fmt.Print("Статусы НЕ 'Отменён':\n")
	for _, status := range orderStatuses {
		if status == "cancelled" {
			continue
		}
		fmt.Printf("%s\n", status)
	}

}
