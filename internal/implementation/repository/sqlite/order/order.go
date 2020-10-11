package orderrepo

import (
	"database/sql"

	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/aeramu/yumfood-go/internal/service/order"
)

// NewRepository initiate repository implementation
func NewRepository() order.Repository {
	db, err := sql.Open("sqlite3", "./internal/database/test.sqlite")
	if err != nil {
		panic(err)
	}
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (r *repository) Put(id string, items []domain.OrderItem) error {
	statement, err := r.db.Prepare("INSERT INTO order_items(order_id, dish_id, vendor_id, amount, request) values(?,?,?,?,?)")
	if err != nil {
		return err
	}
	for _, item := range items {
		_, err = statement.Exec(id, item.DishID, item.VendorID, item.Amount, item.Request)
		if err != nil {
			return err
		}
	}
	statement, err = r.db.Prepare("INSERT INTO orders(id) values(?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAll() ([]domain.Order, error) {
	rows, err := r.db.Query(`SELECT orders.id, items.dish_id, items.vendor_id, items.amount, items.request 
		FROM orders inner join order_items items on orders.id = items.order_id`)
	if err != nil {
		return []domain.Order{}, err
	}
	orders := []domain.Order{}
	items := []domain.OrderItem{}
	lastID := ""
	for rows.Next() {
		item := domain.OrderItem{}
		var id string
		rows.Scan(&id, &item.DishID, &item.VendorID, &item.Amount, &item.Request)
		if lastID != "" && id != lastID {
			orders = append(orders, domain.Order{id, items})
			items = []domain.OrderItem{}
		}
		items = append(items, item)
		lastID = id
	}
	if lastID != "" {
		orders = append(orders, domain.Order{lastID, items})
	}

	return orders, nil
}
