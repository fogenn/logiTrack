package main

import "fmt"

func main() {

	var (
		orderID      int    = 1
		customerName string = "qwertyuiop"
		isDelivered  bool   = false
	)

	orderIDs := []int{}
	orderIDs = append(
		orderIDs,
		0,
		1,
		2,
	)

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
}
