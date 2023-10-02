package mysql

/*
https://gorm.io/zh_CN/docs/query.html
*/

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID      uint `gorm:"primary_key"`
	Name    string
	Email   string
	Balance int64
}

func (User) TableName() string {
	return "users"
}

type Query struct {
	db *gorm.DB
}

func (q *Query) UsersByName(name string) ([]User, error) {
	var users []User
	/*
		SELECT * FROM testdb.users
			WHERE name = "joy"
	*/
	if err := q.db.Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("query UsersByName db.Where err: %w", err)
	}
	return users, nil
}
