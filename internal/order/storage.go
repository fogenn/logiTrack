package order

//type OrderStorage struct {
//	mu     sync.Mutex
//	Orders []Order
//}

//type OrderStorageIntf interface {
//	Save(order *Order)
//	GetAll() []Order
//	GetByID(id int) (Order, error)
//}

type StorageIntf interface {
	Save(order *Order) error
	GetAll() []Order
	GetByID(id int) (*Order, int, error)
	Update(id int, status string) error
	Delete(id int) error
}

//func NewOrderStorage() *OrderStorage {
//	return &OrderStorage{Orders: []Order{}}
//}
//
//// in storage_mock
//func (Os *OrderStorage) Save(order *Order) {
//	if order.ID == 0 {
//		panic("test panic")
//	} else {
//		Os.Orders = append(Os.Orders, *order)
//	}
//	defer fmt.Println("запрос завершён", order)
//}
//
//// in storage_mock
//func (Os OrderStorage) GetAll() []Order {
//	return Os.Orders
//}
//
//// in storage_mock
//func (Os OrderStorage) GetByID(id int) (Order, error) {
//	Orders := Os.Orders
//	for _, order := range Orders {
//		if order.ID == id {
//			return order, nil
//		}
//	}
//	//return Order{}, errors.New("GetByID-order not found")
//	return Order{}, fmt.Errorf("GetByID-order not found %v", Orders)
//}
