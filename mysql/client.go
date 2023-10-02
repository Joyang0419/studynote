package mysql

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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

func GormMySqlFactory(username, password, network, ip, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, ip, port, database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
