package auth

import (
	"github.com/gofrs/uuid"
	"github.com/h1ght1me/auth-micro/pkg/permissions"
)

const identityKey = "id"

type JWTData struct {
	ID     uuid.UUID              `json:"id"`
	Name   string                 `json:"name"`
	Access permissions.Permission `json:"access"`
}
