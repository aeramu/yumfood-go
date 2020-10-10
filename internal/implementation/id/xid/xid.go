package xid

import "github.com/rs/xid"

//NewIDGenerator return id generator implementation
func NewIDGenerator() Generator {
	return Generator{}
}

// Generator empty struct
type Generator struct{}

// Generate a uuid
func (g Generator) Generate() string {
	return xid.New().String()
}
