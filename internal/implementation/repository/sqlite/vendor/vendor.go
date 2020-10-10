package vendor

import (
	"database/sql"

	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/aeramu/yumfood-go/internal/service/vendor"
	_ "github.com/mattn/go-sqlite3" //sqlite 3
)

// NewRepository initiate repository implementation
func NewRepository() vendor.Repository {
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

func (r *repository) Put(id, name, description string) error {
	statement, err := r.db.Prepare("INSERT INTO vendor(id, name, description) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id, name, description string) error {
	statement, err := r.db.Prepare("update vendor set name=?, description=? where id=?")
	_, err = statement.Exec(name, description, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(id string) (domain.Vendor, error) {
	rows, err := r.db.Query("SELECT * FROM vendor where id=?", id)
	if err != nil {
		return domain.Vendor{}, err
	}
	vendor := domain.Vendor{}
	for rows.Next() {
		rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description)
	}
	return vendor, nil
}

func (r *repository) Delete(id string) error {
	statement, err := r.db.Prepare("delete from vendor where id=?")
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
