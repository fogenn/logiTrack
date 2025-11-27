package order

import (
	"fmt"
	"logiTrack/internal/logger"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresOrderStorage struct {
	mu sync.Mutex
	//rwm sync.RWMutex
	//sync
	db *sqlx.DB
}

var _ StorageIntf = (*PostgresOrderStorage)(nil)

func NewPostgresOrderStorage(db *sqlx.DB) *PostgresOrderStorage {
	if err := checkOrdersTableExists(db); err != nil {
		logger.Log.WithError(err).Warn("Таблица orders не существует или недоступна")
	}

	return &PostgresOrderStorage{db: db}
}

func checkOrdersTableExists(db *sqlx.DB) error {
	var exists bool

	query := `
			SELECT EXISTS (
				SELECT FROM information_schema.tables 
				WHERE table_schema = 'public' 
				AND table_name = 'orders'
			);`

	if err := db.QueryRow(query).Scan(&exists); err != nil {
		return fmt.Errorf("Ошибка проверки таблицы 'orders': %w", err)
	} else if !exists {
		return fmt.Errorf("таблица 'orders' не существует")
	}
	return nil
}

func (p *PostgresOrderStorage) Save(order *Order) error {
	//p.mu.Lock()
	//defer p.mu.Unlock()
	defer fmt.Println("запрос завершён", order)

	if order.ID == 0 {
		panic("test panic")
	} else {
		query := `INSERT INTO orders (customer_name, status) VALUES ($1, $2) RETURNING id`
		return p.db.QueryRow(query, order.CustomerName, order.Status).Scan(&order.ID)
	}
	return nil
}

func (p *PostgresOrderStorage) GetAll() []Order {
	//p.mu.Lock()
	//defer p.mu.Unlock()
	query := `SELECT id, customer_name, status FROM orders`
	var orders []Order

	if err := p.db.Select(&orders, query); err != nil {
		logger.Log.WithError(err).Error("Ошибка при выполнении запроса GetAll")
		return nil
	}
	logger.Log.WithField("count", len(orders)).Info("Заказы получены")
	return orders
}

func (p *PostgresOrderStorage) GetByID(id int) (*Order, int, error) {
	//p.mu.Lock()
	//defer p.mu.Unlock()

	var order Order

	query := `SELECT id, customer_name, status FROM orders WHERE id = $1`
	err := p.db.Get(&order, query, id)
	if err != nil {
		return nil, -1, err
	}

	return &order, id, nil
}

// TODO добавить ошибки
func (p *PostgresOrderStorage) Update(id int, status string) error {
	//p.mu.Lock()
	//defer p.mu.Unlock()

	_, _, err := p.GetByID(id)
	if err != nil {
		return err
	}
	query := `UPDATE orders SET status = $1 WHERE id = $2`

	if _, err = p.db.Exec(query, status, id); err != nil {
		return fmt.Errorf("ошибка обновления заказа: %w", err)
	}

	return nil
}

//TODO добавить ошибки на отсутствующий айдишник

func (p *PostgresOrderStorage) Delete(id int) error {
	//p.mu.Lock()
	//defer p.mu.Unlock()

	query := `DELETE FROM orders WHERE id = $1`
	_, err := p.db.Exec(query, id)
	return err
}
