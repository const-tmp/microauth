package database

import (
	"fmt"
	"github.com/h1ght1me/auth-micro/config"
	"github.com/h1ght1me/auth-micro/pkg/database/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(conf *config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		conf.Database.Host,
		conf.Database.DBUser,
		conf.Database.DBPassword,
		conf.Database.DBName,
		conf.Database.Port,
		conf.Database.Timezone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{})
}
