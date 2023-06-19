package db

import (
	"net"
	"test_di/config"
)

type DB struct {
	Conn *net.Conn
}

func ConnectDatabase(c *config.Config) (*DB, error) {
	return &DB{
		Conn: nil,
	}, nil
}
