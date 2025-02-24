package internalsql

import (
	"database/sql"
	_ "database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("Failed to Connect With Databases: %s", err.Error())
		return nil, err
	}
	return db, nil
}
