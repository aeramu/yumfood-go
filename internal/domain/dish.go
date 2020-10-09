package domain

// Dish entity
type Dish struct {
	ID       string
	VendorID string
	Name     string
	Price    int
}

// DishConstructor constructor for dish
type DishConstructor struct {
	ID       string
	VendorID string
	Name     string
	Price    int
}

// New return dish from constructor
func (c DishConstructor) New() Dish {
	return Dish{
		ID:       c.ID,
		VendorID: c.VendorID,
		Name:     c.Name,
		Price:    c.Price,
	}
}
