package twitcasting

import (
	"github.com/rock8526652/DDBOT/lsp/concern"
)

type GroupConcernConfig struct {
	concern.IConfig
}

func (g *GroupConcernConfig) ShouldSendHook(notify concern.Notify) *concern.HookResult {
	return concern.HookResultPass
}

func NewGroupConcernConfig(g concern.IConfig) *GroupConcernConfig {
	return &GroupConcernConfig{g}
}
