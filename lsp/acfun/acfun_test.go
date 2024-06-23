package acfun

import (
	"testing"

	"github.com/rock8526652/DDBOT/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestAcfun(t *testing.T) {
	assert.NotEmpty(t, APath(PathApiChannelList))
	assert.NotEmpty(t, APath("api/channel/list"))
	assert.NotEmpty(t, LiveUrl(test.UID1))
}
