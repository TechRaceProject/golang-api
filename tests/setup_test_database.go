package tests

import (
	services "api/pkg/httputils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the test database")
	}

	services.SetConnection(db)

	return db
}

func GetTestDBConnection() *gorm.DB {
	if db == nil {
		db = setupTestDB()
	}

	return db
}
