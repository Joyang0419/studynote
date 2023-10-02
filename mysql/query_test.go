package mysql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrateUsersTable(t *testing.T) {
	db, err := GormMySqlFactory(
		"root",      /* username */
		"root",      /* password  */
		"tcp",       /* network */
		"localhost", /* ip */
		"3306",      /* port */
		"testdb",    /* database */
	)
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&User{}))
}

func TestQuery_UsersByName(t *testing.T) {
	db, err := GormMySqlFactory(
		"root",      /* username */
		"root",      /* password  */
		"tcp",       /* network */
		"localhost", /* ip */
		"3306",      /* port */
		"testdb",    /* database */
	)
	assert.NoError(t, err)
	q := Query{db: db}
	users, err := q.UsersByName("joy")
	assert.NoError(t, err)
	fmt.Println(users)
}
