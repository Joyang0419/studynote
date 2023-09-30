package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientFactory(t *testing.T) {
	db, err := ClientFactory(
		"root",      /* username */
		"root",      /* password  */
		"tcp",       /* network */
		"localhost", /* ip */
		"3306",      /* port */
		"testdb",    /* database */
	)
	assert.NoError(t, err)
	assert.NoError(t, db.Ping())
}
