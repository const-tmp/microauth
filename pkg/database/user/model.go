package user

import (
	"github.com/gofrs/uuid"
	"github.com/h1ght1me/auth-micro/pkg/permissions"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid"`
	Name      string    `gorm:"size:32"`
	Password  []byte    `gorm:"size:60"`
	Access    permissions.Permission
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, err = uuid.NewV4()
	return
}
