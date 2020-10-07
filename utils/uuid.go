package utils


import (
	"github.com/google/uuid"
)

// UUID alias
type UUID = uuid.UUID

// create uuid
func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

// return uuid
func getUUID() UUID {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}


