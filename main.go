package main

import (
	"fmt"
	"strings"
)

type OrderStorage interface {
	Save(order *Order)
	GetAll() []Order
}

type Order struct {
	ID           int
	CustomerName string
	Status       string
}

type OrderStorageMock struct {
	orders []Order
}

func (o Order) IsDelivered() bool {
	return o.Status == "Delivered"
}

func (OM *OrderStorageMock) Save(order *Order) {
	OM.orders = append(OM.orders, *order)
}

func (OM OrderStorageMock) GetAll() []Order {
	return OM.orders
}

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

	OrderStorageMock := OrderStorageMock{}

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

	Order1 := Order{
		ID:           2,
		CustomerName: "Ivan",
		Status:       "Delivered",
	}
	Order2 := Order{
		ID:           3,
		CustomerName: "Aleksandr",
		Status:       "Cancelled",
	}

	OrderStorageMock.Save(&Order1)
	OrderStorageMock.Save(&Order2)

	fmt.Println(OrderStorageMock.GetAll())

	Order2 = Order{
		ID:           3,
		CustomerName: "Maxim",
		Status:       "Cancelled",
	}
	fmt.Println(OrderStorageMock.GetAll())

}
