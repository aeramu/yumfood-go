package domain

// Order entity
type Order struct {
	ID    string      `json:"id"`
	Items []OrderItem `json:"items"`
}

// OrderItem entity inside Order entity
type OrderItem struct {
	DishID  string `json:"dish_id"`
	Amount  int    `json:"amount"`
	Request string `json:"request"`
}
