package uuid

import "github.com/google/uuid"

//NewIDGenerator return id generator implementation
func NewIDGenerator() Generator {
	return Generator{}
}

// Generator empty struct
type Generator struct{}

// Generate a uuid
func (g Generator) Generate() string {
	return uuid.New().String()
}
