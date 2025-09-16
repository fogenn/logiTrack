package main

import (
	"fmt"
	"strings"
)

func safeFuncSaveOrder(order *Order, storage *OrderStorageMock) { // чисто эксперемент
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Перехвачена паника:", r)
		}
	}()
	storage.Save(order)
}

type OrderStorage interface {
	Save(order *Order)
	GetAll() []Order
	GetByID(id int) (Order, error)
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

func (Osm *OrderStorageMock) Save(order *Order) {
	if order.ID == 0 {
		panic("test panic")
	} else {
		Osm.orders = append(Osm.orders, *order)
	}
	defer fmt.Println("запрос завершён", order)
}

func (Osm OrderStorageMock) GetAll() []Order {
	return Osm.orders
}

func (Osm OrderStorageMock) GetByID(id int) (Order, error) {
	Orders := Osm.orders
	for _, order := range Orders {
		if order.ID == id {
			return order, nil
		}
	}
	//return Order{}, errors.New("GetByID-order not found")
	return Order{}, fmt.Errorf("GetByID-order not found %v", Orders)
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

func TESTSafeFuncGetByID(ById int, storage OrderStorageMock) {
	order, err := storage.GetByID(ById)
	if err != nil {
		fmt.Println(fmt.Errorf("TESTSafeFuncGetByID-Не удалось найти заказ  %+v: %w", order, err))

	} else {
		fmt.Printf("Корректный ID - %v", order)
	}
}

func main() {
	//safeFunc()
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Перехвачена паника:", r)
	//	}
	//}()

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
	Order3 := Order{
		ID:           0,
		CustomerName: "",
		Status:       "",
	}
	Order4 := Order{
		ID:           4,
		CustomerName: "Norm",
		Status:       "Delivered",
	}

	OrderStorageMock.Save(&Order1)
	safeFuncSaveOrder(&Order3, &OrderStorageMock) // попытка записать опасный слайс
	OrderStorageMock.Save(&Order2)
	safeFuncSaveOrder(&Order4, &OrderStorageMock) // попытка записать норм слайс в безопасной функции

	fmt.Println(OrderStorageMock.GetAll())
	fmt.Println("-----------------------")
	fmt.Println(OrderStorageMock.orders[0].ID)
	fmt.Println("----------Здесь вызывается сам метод .GetByID с правильным ID-------------")
	order, err := OrderStorageMock.GetByID(2)
	if err != nil {
		fmt.Println(fmt.Errorf("Не удалось найти заказ  %v: %w", order, err))
	} else {
		fmt.Printf("Корректный ID - %v\n", order)
	}
	fmt.Println("----------Здесь вызывается сам метод .GetByID с НЕправильным ID-------------")
	order1, err1 := OrderStorageMock.GetByID(1)
	if err1 != nil {
		fmt.Println(fmt.Errorf("Не удалось найти заказ  %v: %w", order1, err1))
	} else {
		fmt.Printf("Корректный ID - %v\n", order1)
	}

	fmt.Println("+++++++++++++++++++")
	fmt.Println(OrderStorageMock.GetByID(2))
	fmt.Println("-----------Здесь вызывается TESTSafeFuncGetByID с НЕправильным ID------------")
	TESTSafeFuncGetByID(1, OrderStorageMock)
	fmt.Println("-----------Здесь вызывается TESTSafeFuncGetByID с правильным ID------------")
	TESTSafeFuncGetByID(4, OrderStorageMock)

}
