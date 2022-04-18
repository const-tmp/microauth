package web

import (
	"github.com/gofrs/uuid"
	"github.com/h1ght1me/auth-micro/pkg/permissions"
)

type APIResponse struct {
	OK     bool        `json:"ok"`
	Result interface{} `json:"result,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type Auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type JWTData struct {
	ID     uuid.UUID               `json:"id"`
	Name   string                  `json:"name"`
	Access permissions.Permissions `json:"access"`
}
