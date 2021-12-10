package impl

import (
	"testing"

	"github.com/mniak/pismo/internal/abstractions"
	"github.com/stretchr/testify/assert"
)

func TestImplementation(t *testing.T) {
	assert.Implements(t, (*abstractions.PasswordManager)(nil), new(KeepassPasswordManager))
}
