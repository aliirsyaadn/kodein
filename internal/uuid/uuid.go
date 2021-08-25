package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

func New() string {
	return fmt.Sprint(uuid.New())
}

func NewV1() string {
	raw, _ := uuid.NewUUID()
	return fmt.Sprint(raw)
}

func IsValid(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	}
	return true
}
