package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func safeFuncSaveOrder(order *Order, storage *OrderStorageMock) { // чисто эксперемент
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Перехвачена паника:", r)
		}
	}()
	storage.Save(order)
}

// in model
type OrderStorage interface {
	Save(order *Order)
	GetAll() []Order
	GetByID(id int) (Order, error)
}

// in model
type Order struct {
	ID           int
	CustomerName string
	Status       string
}

type OrderStorageMock struct {
	orders []Order
}

// TODO надо наверно избавиться
func (o Order) IsDelivered() bool {

	return o.Status == "Delivered"
}

// in storage_mock
func (Osm *OrderStorageMock) Save(order *Order) {
	if order.ID == 0 {
		panic("test panic")
	} else {
		Osm.orders = append(Osm.orders, *order)
	}
	defer fmt.Println("запрос завершён", order)
}

// in storage_mock
func (Osm OrderStorageMock) GetAll() []Order {
	return Osm.orders
}

// in storage_mock
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

// TODO надо наверно избавиться
func markCancelled(status *string) []string {
	if *status == "created" {
		*status = "cancelled"
	}
	return nil
}

// in Service
func TESTSafeGetByID(ById int, storage OrderStorageMock) {
	order, err := storage.GetByID(ById)
	if err != nil {
		fmt.Println(fmt.Errorf("TESTSafeGetByID-Не удалось найти заказ  %+v: %w", order, err))

	} else {
		fmt.Printf("Корректный ID - %v", order)
	}
}

// in Service
func StartDeliveryWorker(ch chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case order, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("\nOrder ID %v in worke -", order.ID)
			time.Sleep(3 * time.Second)
			fmt.Printf("\n - Order ID %v delivered", order.ID)

		case <-time.After(1 * time.Second):
			fmt.Println("\nДолго ждём, подождём ещё...")
		}

	}
	//TODO надо наверно избавиться
	//for order := range ch {
	//	select {
	//	case <-time.After(1 * time.Second):
	//		fmt.Println("Долго ждём, скип...")
	//	default:
	//		fmt.Printf("\nOrder ID %v in worke -", order.ID)
	//		time.Sleep(3 * time.Second)
	//		fmt.Printf("\n - Order ID %v delivered", order.ID)
	//	}
	//}
}

func main() {

	//TODO надо наверно избавиться
	var (
		orderID      int    = 1
		customerName string = "ИВАН"
		isDelivered  bool   = false
	)
	//TODO надо наверно избавиться
	orderIDs := []int{}
	orderIDs = append(
		orderIDs,
		0,
		1,
		2,
	)

	OrderStorageMock := OrderStorageMock{}

	//TODO надо наверно избавиться
	orderStatuses := []string{
		"created", "shipped", "delivered", "cancelled", "delivered", "created",
	}

	//TODO надо наверно избавиться
	orderCount := map[string]int{
		"Client1": 1,
		"Client2": 5,
	}

	//TODO надо наверно избавиться
	isReadyToShip := orderCount["Client2"] > 2

	//TODO надо наверно избавиться
	fmt.Println(
		orderID,
		customerName,
		isDelivered,
	)

	//TODO надо наверно избавиться
	fmt.Println(
		len(orderIDs),
		cap(orderIDs),
	)

	//TODO надо наверно избавиться
	fmt.Println(isReadyToShip)

	//TODO надо наверно избавиться
	countDelivered(orderStatuses)

	//TODO надо наверно избавиться
	fmt.Println("слайс статусов ДО отмены -", orderStatuses)

	//TODO надо наверно избавиться
	for i := range orderStatuses {
		markCancelled(&orderStatuses[i])
	}

	//TODO надо наверно избавиться
	fmt.Println("слайс статусов после отмены -", orderStatuses, "\n")

	//TODO надо наверно избавиться
	fmt.Print("Статусы НЕ 'Отменён':\n")
	for _, status := range orderStatuses {
		if status == "cancelled" {
			continue
		}
		fmt.Printf("%s\n", status)
	}

	//Order1 := Order{
	//	ID:           2,
	//	CustomerName: "Ivan",
	//	Status:       "Delivered",
	//}
	//Order2 := Order{
	//	ID:           3,
	//	CustomerName: "Aleksandr",
	//	Status:       "Cancelled",
	//}
	//Order3 := Order{
	//	ID:           0,
	//	CustomerName: "",
	//	Status:       "",
	//}
	//Order4 := Order{
	//	ID:           4,
	//	CustomerName: "Norm",
	//	Status:       "Delivered",
	//}

	//Order5 :=[]Order{
	//	{
	//	ID:           2,
	//		CustomerName: "Ivan",
	//		Status:       "Delivered",
	//	},
	//	{
	//	ID:           2,
	//		CustomerName: "Ivan",
	//		Status:       "Delivered",
	//	},
	//	{
	//	ID:           2,
	//		CustomerName: "Ivan",
	//		Status:       "Delivered",
	//	},
	//	{
	//	ID:           2,
	//		CustomerName: "Ivan",
	//		Status:       "Delivered",
	//	},
	//}

	Order1 := Order{
		ID:           1,
		CustomerName: "Ivan",
		Status:       "Delivered",
	}
	Order2 := Order{
		ID:           2,
		CustomerName: "Aleksandr",
		Status:       "Cancelled",
	}
	Order3 := Order{
		ID:           3,
		CustomerName: "Maria",
		Status:       "Pending",
	}
	Order4 := Order{
		ID:           4,
		CustomerName: "Norm",
		Status:       "Shipped",
	}
	Order5 := Order{
		ID:           5,
		CustomerName: "Elena",
		Status:       "Processing",
	}
	Order6 := Order{
		ID:           6,
		CustomerName: "Dmitry",
		Status:       "Delivered",
	}
	Order7 := Order{
		ID:           7,
		CustomerName: "Olga",
		Status:       "Returned",
	}
	Order8 := Order{
		ID:           8,
		CustomerName: "Sergey",
		Status:       "Awaiting Payment",
	}
	Order9 := Order{
		ID:           9,
		CustomerName: "Tatiana",
		Status:       "Cancelled",
	}

	OrderStorageMock.Save(&Order1)
	safeFuncSaveOrder(&Order3, &OrderStorageMock) // попытка записать опасный слайс
	OrderStorageMock.Save(&Order2)
	safeFuncSaveOrder(&Order4, &OrderStorageMock) // попытка записать норм слайс в безопасной функции
	safeFuncSaveOrder(&Order5, &OrderStorageMock)
	safeFuncSaveOrder(&Order6, &OrderStorageMock)
	safeFuncSaveOrder(&Order7, &OrderStorageMock)
	safeFuncSaveOrder(&Order8, &OrderStorageMock)
	safeFuncSaveOrder(&Order9, &OrderStorageMock)

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
	fmt.Println("-----------Здесь вызывается TESTSafeGetByID с НЕправильным ID------------")
	TESTSafeGetByID(1, OrderStorageMock)
	fmt.Println("-----------Здесь вызывается TESTSafeGetByID с правильным ID------------")
	TESTSafeGetByID(4, OrderStorageMock)
	fmt.Println("\n+++++++++++++++++++")
	fmt.Println("+++++++++++++++++++")

	fmt.Println("-----------Здесь вызывается StartDeliveryWorker которая  вычитывает из канала------------")
	deliveryChan := make(chan Order, 2)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go StartDeliveryWorker(deliveryChan, &wg)

	for _, orderU := range OrderStorageMock.orders {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Printf("\nТо что в канал положили: %+v", orderU)
		deliveryChan <- orderU
	}
	close(deliveryChan)

	wg.Wait()
	fmt.Println("\nВсе заказы выполнены")
}
