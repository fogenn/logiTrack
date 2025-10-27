package order

type Order struct {
	ID           int
	CustomerName string
	Status       string
}

type OrderStorageIntf interface {
	Save(order *Order)
	GetAll() []Order
	GetByID(id int) (Order, error)
}
