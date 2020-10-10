package shortid

import "github.com/ventu-io/go-shortid"

//NewUIDGenerator return uid generator implementation
func NewUIDGenerator() Generator {
	return Generator{}
}

// Generator empty struct
type Generator struct{}

// Generate a uuid
func (g Generator) Generate() string {
	return shortid.MustGenerate()
}
