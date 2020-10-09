package domain

// Order entity
type Order struct {
	ID    string
	Items []OrderItem
}

// OrderItem entity inside Order entity
type OrderItem struct {
	DishID  string
	Amount  int
	Request string
}
