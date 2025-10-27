package order

import "fmt"

type OrderStorage struct {
	orders []Order
}

func NewOrderStorage(os []Order) *OrderStorage {
	return &OrderStorage{orders: os}
}

// in storage_mock
func (Os *OrderStorage) Save(order *Order) {
	if order.ID == 0 {
		panic("test panic")
	} else {
		Os.orders = append(Os.orders, *order)
	}
	defer fmt.Println("запрос завершён", order)
}

// in storage_mock
func (Os OrderStorage) GetAll() []Order {
	return Os.orders
}

// in storage_mock
func (Os OrderStorage) GetByID(id int) (Order, error) {
	Orders := Os.orders
	for _, order := range Orders {
		if order.ID == id {
			return order, nil
		}
	}
	//return Order{}, errors.New("GetByID-order not found")
	return Order{}, fmt.Errorf("GetByID-order not found %v", Orders)
}
