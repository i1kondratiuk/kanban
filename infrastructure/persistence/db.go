package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// NewDbConnection ...
func NewDbConnection(host, port, user, password, dbname, driver string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open up our database connection.
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
