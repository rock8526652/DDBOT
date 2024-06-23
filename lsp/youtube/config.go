package youtube

import (
	"github.com/rock8526652/DDBOT/lsp/concern"
)

type GroupConcernConfig struct {
	concern.IConfig
}

func (g *GroupConcernConfig) ShouldSendHook(notify concern.Notify) *concern.HookResult {
	if c, ok := notify.(*ConcernNotify); ok {
		// 直播预告也应该推送
		if c.IsWaiting() {
			return concern.HookResultPass
		}
	}
	return g.IConfig.ShouldSendHook(notify)
}

func NewGroupConcernConfig(g concern.IConfig) *GroupConcernConfig {
	return &GroupConcernConfig{g}
}
