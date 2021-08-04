package internal

import (
	"fmt"
)

var DBInstance  DB

// DB implements a generic interface for DB clients
type DB interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string)
	GetForColumnValue(field string, value interface{}) ([]string, error)
}

// NewDB initializes the cache instance based on Config
func NewDB(dbType string) error {
	var err error

	switch dbType {
	case "memory":
		DBInstance = NewDBSchema()

	default:
		err = fmt.Errorf("unknown database type: %s", dbType)
	}

	return err
}