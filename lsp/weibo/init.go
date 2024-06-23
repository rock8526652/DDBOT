package weibo

import (
	"net/http"
	"time"

	"github.com/rock8526652/DDBOT/lsp/concern"
	"github.com/rock8526652/DDBOT/requests"
	localutils "github.com/rock8526652/DDBOT/utils"
)

func init() {
	concern.RegisterConcern(NewConcern(concern.GetNotifyChan()))

	var cookies []*http.Cookie
	var err error

	freshCookieOpt := func() {
		localutils.Retry(3, time.Second, func() bool {
			cookies, err = FreshCookie()
			return err == nil
		})
		if err != nil {
			logger.Errorf("FreshCookie error %v", err)
		} else {
			var opt []requests.Option
			for _, cookie := range cookies {
				opt = append(opt, requests.HttpCookieOption(cookie))
			}
			visitorCookiesOpt.Store(opt)
		}
	}
	freshCookieOpt()
	go func() {
		for range time.Tick(time.Hour) {
			freshCookieOpt()
		}
	}()
}
