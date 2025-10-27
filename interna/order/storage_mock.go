package order

import "fmt"

type OrderStorageMock struct {
	orders []Order
}

func NewOrderStorageMock(os []Order) *OrderStorageMock {
	return &OrderStorageMock{orders: os}
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
