package bilibili

import (
	"testing"

	"github.com/rock8526652/DDBOT/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestRelationModify(t *testing.T) {
	_, err := RelationModify(test.UID1, ActSub)
	assert.NotNil(t, err)
}
