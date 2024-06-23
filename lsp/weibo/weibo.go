package weibo

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/rock8526652/DDBOT/requests"
	"go.uber.org/atomic"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	Site = "weibo"
)

var (
	visitorCookiesOpt atomic.Value
)

func CookieOption() []requests.Option {
	if c := visitorCookiesOpt.Load(); c != nil {
		return c.([]requests.Option)
	}
	return nil
}
