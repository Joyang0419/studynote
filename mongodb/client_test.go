package mongodb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientFactory(t *testing.T) {
	mongodb, err := ClientFactory(
		context.TODO(),
		"root",      /* username */
		"root",      /* password */
		"localhost", /* ip */
		"27017",     /* port */
	)
	assert.NoError(t, err)
	assert.NoError(t, mongodb.Ping(context.TODO(), nil))
}
