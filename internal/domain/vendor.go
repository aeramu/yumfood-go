package domain

// Vendor entity
type Vendor struct {
	ID          string
	Name        string
	Description string
}

// VendorConstructor constructor for vendor
type VendorConstructor struct {
	ID          string
	Name        string
	Description string
}

// New return vendor from constructor
func (c VendorConstructor) New() Vendor {
	if len(c.Name) >= 128 {
		c.Name = c.Name[:127]
	}
	return Vendor{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
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
