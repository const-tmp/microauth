package database

import (
	"github.com/h1ght1me/auth-micro/config"
	"testing"
)

func TestDBConn(t *testing.T) {
	cfg, err := config.LoadConfig("test")
	if err != nil {
		t.Error(err)
	}

	_, err = Connect(cfg)
	if err != nil {
		t.Error(err)
	}
}

func TestMigration(t *testing.T) {
	cfg, err := config.LoadConfig("test")
	if err != nil {
		t.Error(err)
	}

	db, err := Connect(cfg)
	if err != nil {
		t.Error(err)
	}

	err = RunMigration(db)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUSer(t *testing.T) {
	cfg, err := config.LoadConfig("test")
	if err != nil {
		t.Error(err)
	}

	db, err := Connect(cfg)
	if err != nil {
		t.Error(err)
	}

	err = RunMigration(db)
	if err != nil {
		t.Error(err)
	}
}
