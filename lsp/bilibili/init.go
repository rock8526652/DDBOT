package bilibili

import (
	"time"

	"github.com/rock8526652/DDBOT/lsp/concern"
)

func init() {
	concern.RegisterConcern(NewConcern(concern.GetNotifyChan()))
	refreshCookieJar()
	refreshNavWbi()
	go func() {
		for range time.Tick(time.Minute * 60) {
			refreshCookieJar()
		}
	}()
	go func() {
		for range time.Tick(2 * time.Minute) {
			refreshNavWbi()
		}
	}()
}
