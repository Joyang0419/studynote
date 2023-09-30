package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ClientFactory(username, password, network, ip, port, database string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, ip, port, database)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, fmt.Errorf("ClientFactory sql.Open err: %w", err)
	}

	return db, nil
}
