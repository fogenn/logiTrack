package order

import (
	"fmt"
	"strings"
	"sync"
)

type OrderStorageMock struct {
	mu     sync.Mutex
	Orders []Order
}

type OrderStorageMockIntf interface {
	Save(order *Order)
	GetAll() []Order
	GetByID(id int) (*Order, int, error)
	Update(id int, status string) error
	Delete(id int) error
}

func NewOrderStorageMock() *OrderStorageMock {
	return &OrderStorageMock{Orders: []Order{}}
}

// in storage_mock
func (osm *OrderStorageMock) Save(order *Order) {
	osm.mu.Lock()
	defer osm.mu.Unlock()

	if order.ID == 0 {
		panic("test panic")
	} else {
		osm.Orders = append(osm.Orders, *order)
	}
	defer fmt.Println("запрос завершён", order)
}

// in storage_mock
func (osm OrderStorageMock) GetAll() []Order {
	osm.mu.Lock()
	defer osm.mu.Unlock()

	return osm.Orders
}

// in storage_mock
func (osm *OrderStorageMock) GetByID(id int) (*Order, int, error) {
	osm.mu.Lock()
	defer osm.mu.Unlock()

	for i, order := range osm.Orders {
		if order.ID == id {
			return &order, i, nil
		}
	}
	return &Order{}, -1, fmt.Errorf("GetByID-order not found %v", osm.Orders)
}

func (osm *OrderStorageMock) Update(id int, status string) error {
	osm.mu.Lock()
	defer osm.mu.Unlock()

	_, i, err := osm.GetByID(id)
	if err != nil {
		return err
	}

	switch strings.ToLower(status) {
	case "shipped":
		osm.Orders[i].Status = status
	case "delivered":
		osm.Orders[i].Status = status
	case "cancelled":
		osm.Orders[i].Status = status
	case "processing":
		osm.Orders[i].Status = status
	default:
		return fmt.Errorf("invalid status %v", status)
	}
	return nil
}

func (osm *OrderStorageMock) Delete(id int) error {
	osm.mu.Lock()
	defer osm.mu.Unlock()

	_, index, err := osm.GetByID(id)
	lastIndex := len(osm.Orders) - 1
	if err != nil {
		return err
	}
	osm.Orders[index] = osm.Orders[lastIndex]
	osm.Orders = osm.Orders[:lastIndex]
	return nil
}
