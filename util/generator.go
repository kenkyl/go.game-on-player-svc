package util

import (
	"github.com/google/uuid"
)

const available = 999999999
const prefix = "PLAYER@"

func GeneratePlayerID() string {
	id := uuid.New()
	str := prefix + id.String()
	return str
}
