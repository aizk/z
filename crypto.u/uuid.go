package crypto_u

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func GetUUIDWithoutRail() string {
	return strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", -1)
}
