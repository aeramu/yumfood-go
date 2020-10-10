package domain

// Dish entity
type Dish struct {
	ID       string `json:"id"`
	VendorID string `json:"vendor_id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}
