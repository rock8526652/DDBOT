package version

import (
	"testing"

	"github.com/rock8526652/DDBOT/internal/test"
	localdb "github.com/rock8526652/DDBOT/lsp/buntdb"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/buntdb"
)

func TestGetCurrentVersion(t *testing.T) {
	test.InitBuntdb(t)
	defer test.CloseBuntdb(t)

	assert.Zero(t, GetCurrentVersion(test.VersionName))

	var testCase = []int64{
		0, 1, 2, 3, 4, 3,
	}

	var expected = []int64{
		0, 0, 1, 2, 3, 4,
	}

	for idx, v := range testCase {
		old, err := SetVersion(test.VersionName, v)
		assert.Nil(t, err)
		assert.Equal(t, v, GetCurrentVersion(test.VersionName))
		assert.Equal(t, expected[idx], old)
	}

	err := localdb.RWCoverTx(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(localdb.VersionKey(test.VersionName), "wrong", nil)
		return err
	})
	assert.Nil(t, err)
	assert.EqualValues(t, -1, GetCurrentVersion(test.VersionName))
}
