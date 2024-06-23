package acfun

import (
	"testing"
	"time"

	localutils "github.com/rock8526652/DDBOT/utils"
	"github.com/stretchr/testify/assert"
)

func TestApiChannelList(t *testing.T) {
	var resp *ApiChannelListResponse
	var err error
	localutils.Retry(5, time.Second, func() bool {
		resp, err = ApiChannelList(100, "")
		return err == nil
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
