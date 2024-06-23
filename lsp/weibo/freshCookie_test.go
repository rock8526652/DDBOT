package weibo

import (
	"net/http"
	"testing"
	"time"

	localutils "github.com/rock8526652/DDBOT/utils"
	"github.com/stretchr/testify/assert"
)

func TestFreshCookie(t *testing.T) {
	var cookies []*http.Cookie
	var err error
	localutils.Retry(5, time.Second, func() bool {
		cookies, err = FreshCookie()
		return err == nil
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, cookies)
}
