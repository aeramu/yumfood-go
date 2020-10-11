package vendor

import (
	"database/sql"
	"strings"

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
	statement, err := r.db.Prepare("INSERT INTO vendors(id, name, description) values(?,?,?)")
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
	statement, err := r.db.Prepare("update vendors set name=?, description=? where id=?")
	_, err = statement.Exec(name, description, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(id string) (domain.Vendor, error) {
	rows, err := r.db.Query("SELECT * FROM vendors where id=?", id)
	if err != nil {
		return domain.Vendor{}, err
	}
	vendor := domain.Vendor{}
	for rows.Next() {
		rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description)
	}
	return vendor, nil
}

func (r *repository) GetAll() ([]domain.Vendor, error) {
	rows, err := r.db.Query("SELECT * FROM vendors")
	if err != nil {
		return []domain.Vendor{}, err
	}
	vendors := []domain.Vendor{}
	for rows.Next() {
		vendor := domain.Vendor{}
		rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description)
		vendors = append(vendors, vendor)
	}
	return vendors, nil
}

func (r *repository) GetListByTags(tags []string) ([]domain.Vendor, error) {
	args := []interface{}{}
	for _, tag := range tags {
		args = append(args, tag)
	}
	args = append(args, len(tags))
	rows, err := r.db.Query(`
		SELECT vendors.id, vendors.name, vendors.description 
		FROM vendors
		inner join vendor_tags tags on vendors.id = tags.vendor_id
		WHERE tags.tag IN (?`+strings.Repeat(", ?", len(tags)-1)+`)
		GROUP BY vendors.id HAVING COUNT(DISTINCT tags.tag) = ? `, args...)
	if err != nil {
		return []domain.Vendor{}, err
	}
	vendors := []domain.Vendor{}
	for rows.Next() {
		vendor := domain.Vendor{}
		rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description)
		vendors = append(vendors, vendor)
	}
	return vendors, nil
}

func (r *repository) Delete(id string) error {
	statement, err := r.db.Prepare("delete from vendors where id=?")
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Tag(id, tag string) error {
	rows, err := r.db.Query("SELECT * FROM vendor_tags where vendor_id=? and tag=?", id, tag)
	if err != nil {
		return err
	}
	if rows.Next() {
		return nil
	}
	statement, err := r.db.Prepare("INSERT INTO vendor_tags(vendor_id, tag) values(?,?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id, tag)
	if err != nil {
		return err
	}
	return nil
}
