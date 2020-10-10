package vendor

import (
	"database/sql"

	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/aeramu/yumfood-go/internal/service/dish"
)

// NewRepository initiate repository implementation
func NewRepository() dish.Repository {
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

func (r *repository) Put(id, vendorID, name string, price int) error {
	statement, err := r.db.Prepare("INSERT INTO dish(id, vendor_id, name, price) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id, vendorID, name, price)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id, name string, price int) error {
	statement, err := r.db.Prepare("update dish set name=?, price=? where id=?")
	_, err = statement.Exec(name, price, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) Get(id string) (domain.Dish, error) {
	rows, err := r.db.Query("SELECT * FROM dish where id=?", id)
	if err != nil {
		return domain.Dish{}, err
	}
	dish := domain.Dish{}
	for rows.Next() {
		rows.Scan(&dish.ID, &dish.VendorID, &dish.Name, &dish.Price)
	}
	return dish, nil
}

func (r *repository) Delete(id string) error {
	statement, err := r.db.Prepare("delete from dish where id=?")
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetListByVendorID(vendorID string) ([]domain.Dish, error) {
	rows, err := r.db.Query("SELECT * FROM dish where vendor_id=?", vendorID)
	if err != nil {
		return []domain.Dish{}, err
	}
	dishes := []domain.Dish{}
	for rows.Next() {
		dish := domain.Dish{}
		rows.Scan(&dish.ID, &dish.VendorID, &dish.Name, &dish.Price)
		dishes = append(dishes, dish)
	}
	return dishes, nil
}
