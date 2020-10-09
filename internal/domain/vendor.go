package domain

// Vendor entity
type Vendor struct {
	ID          string
	Name        string
	Description string
}

// NewVendor constructor for vendor
func NewVendor(name string, description string) Vendor {
	if len(name) >= 128 {
		name = name[:127]
	}
	return Vendor{
		Name:        name,
		Description: description,
	}
}

// Update change vendor state
func (v *Vendor) Update(name string, description string) *Vendor {
	if len(name) >= 128 {
		name = name[:127]
	}
	v.Name = name
	v.Description = description
	return v
}
