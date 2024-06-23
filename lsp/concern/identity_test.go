package concern

import (
	"testing"

	"github.com/rock8526652/DDBOT/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestNewIdentity(t *testing.T) {
	a := NewIdentity(test.ID1, test.NAME1)
	assert.EqualValues(t, test.ID1, a.GetUid())
	assert.EqualValues(t, test.NAME1, a.GetName())
}
